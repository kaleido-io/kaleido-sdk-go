package registry

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	kaleido "github.com/kaleido-io/kaleido-sdk-go/common"
	"github.com/kaleido-io/kaleido-sdk-go/contracts/directory"
	"github.com/youmark/pkcs8"
	jose "gopkg.in/square/go-jose.v2"
)

// Organization ...
type Organization struct {
	Consortium     string `json:"consortia_id,omitempty"`
	Environment    string `json:"environment_id,omitempty"`
	MemberID       string `json:"membership_id,omitempty"`
	Name           string `json:"-"`
	Owner          string `json:"-"`
	SigningKeyFile string `json:"-"`
	CertPEMFile    string `json:"-"`
}

// VerifiedOrganization ...
type VerifiedOrganization struct {
	ID       string            `json:"id,omitempty"`
	Name     string            `json:"name,omitempty"`
	Owner    string            `json:"owner,omitempty"`
	Proof    *JSONWebSignature `json:"proof,omitempty"`
	ParentID string            `json:"parent,omitempty"`
}

type ContractOrganization struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Owner    string `json:"owner,omitempty"`
	ParentID string `json:"parent,omitempty"`
}

// JSONWebSignature json representation of JWS
type JSONWebSignature struct {
	Headers    []string `json:"headers"`
	Payload    string   `json:"payload"`
	Signatures []string `json:"signatures"`
}

// SignedRequest signed payload for identity creation of an org
type SignedRequest struct {
	Consortium   string           `json:"consortia_id,omitempty"`
	Environment  string           `json:"environment_id,omitempty"`
	MembershipID string           `json:"membership_id,omitempty"`
	JWS          JSONWebSignature `json:"jwsjs,omitempty"`
}

func (org *Organization) generateNonce() (string, error) {
	type responseBody struct {
		Nonce string `json:"nonce,omitempty"`
	}

	client := Utils().getAPIClient()
	var noncePayload responseBody
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		// SetBody(org).
		SetBody(org).
		SetResult(&noncePayload).
		Post("/nonce")

	if err != nil {
		return "", err
	}
	return noncePayload.Nonce, Utils().validateCreateResponse(response, err, "nonce")
}

// sourced from go-ethereum
func zeroKey(k *ecdsa.PrivateKey) {
	b := k.D.Bits()
	for i := range b {
		b[i] = 0
	}
}

func (org *Organization) createSignedRequestForRegistration() (*SignedRequest, error) {
	request := SignedRequest{
		Consortium:   org.Consortium,
		Environment:  org.Environment,
		MembershipID: org.MemberID,
	}

	// read the key file
	pemEncodedBytes, err := ioutil.ReadFile(org.SigningKeyFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemEncodedBytes)
	der := block.Bytes

	var ecdsaKey *ecdsa.PrivateKey
	if strings.Contains(string(pemEncodedBytes), "-----BEGIN ENCRYPTED PRIVATE KEY-----") {
		passphrase, err := Utils().readPassword("KLD_PKCS8_SIGNING_KEY_PASSPHRASE", "Encrypted signing PKCS8 key requires a password:")
		if err != nil {
			return nil, err
		}
		privateKey, err := pkcs8.ParsePKCS8PrivateKey(der, []byte(passphrase))
		if err != nil {
			return nil, err
		}
		ecdsaKey = privateKey.(*ecdsa.PrivateKey)
	} else {
		privateKey, err := pkcs8.ParsePKCS8PrivateKey(der)
		if err != nil {
			return nil, err
		}
		ecdsaKey = privateKey.(*ecdsa.PrivateKey)
	}
	defer zeroKey(ecdsaKey)

	// read the provided proof
	proofPEM, err := ioutil.ReadFile(org.CertPEMFile)
	if err != nil {
		return nil, err
	}

	certBlock, _ := pem.Decode(proofPEM)
	if certBlock == nil {
		return nil, errors.New("failed to parse certificate")
	}
	cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return nil, err
	}

	CNTokens := strings.Split(cert.Subject.CommonName, "-")
	if len(CNTokens) < 4 {
		return nil, errors.New("Certificate common name does not follow the format of <orgid>-<nonce>--<name>")
	}

	type responseBody struct {
		OrgName string `json:"org_name,omitempty"`
	}

	clientNetMgr := Utils().GetNetworkManagerClient()
	targetURL := "/consortia/" + org.Consortium + "/memberships/" + org.MemberID
	var memberPayload responseBody
	memberResponse, err := clientNetMgr.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&memberPayload).
		Get(targetURL)

	err = Utils().ValidateGetResponse(memberResponse, err, "membership")
	if err != nil {
		return nil, err
	}

	registryName := memberPayload.OrgName + "-" + org.MemberID

	// create a new signer using ECDSA (ES256) algorithm with the given private key
	var alg jose.SignatureAlgorithm
	switch ecdsaKey.Curve.Params().BitSize {
	case 256:
		alg = jose.ES256
	case 384:
		alg = jose.ES384
	case 521: // not a typo, ES512 == 521 curve bits
		alg = jose.ES512
	}
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: alg, Key: ecdsaKey}, nil)
	if err != nil {
		return nil, err
	}

	// create the json payload that needs to be signed
	nonce, err := org.generateNonce()
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(map[string]interface{}{
		"envId":   org.Environment,
		"memId":   org.MemberID,
		"nonce":   nonce,
		"name":    registryName,
		"proof":   string(proofPEM),
		"address": org.Owner})

	if err != nil {
		return nil, err
	}

	object, err := signer.Sign(jsonBytes)
	if err != nil {
		return nil, err
	}

	serialized, _ := object.CompactSerialize()
	tokens := strings.Split(serialized, ".")
	request.JWS.Headers = append(request.JWS.Headers, tokens[0])
	request.JWS.Payload = tokens[1]
	request.JWS.Signatures = append(request.JWS.Signatures, tokens[2])

	return &request, nil
}

func (org *Organization) populateServiceTargets() error {
	var service *ServiceDefinitionType
	var err error
	if service, err = Utils().GetServiceDefinition(); err != nil {
		return err
	}
	org.Consortium = service.Consortium
	org.Environment = service.Environment

	return nil
}

// InvokeCreate registers a verified organization with the on-chain registry
// and stores the proof on-chain
func (org *Organization) InvokeCreate() (*VerifiedOrganization, error) {
	// if consortium or environment is not set, retrieve it from the service definition
	if org.Consortium == "" || org.Environment == "" {
		if err := org.populateServiceTargets(); err != nil {
			return nil, err
		}
	}

	// sign payload
	signedPayload, err := org.createSignedRequestForRegistration()
	if err != nil {
		return nil, err
	}

	client := Utils().getAPIClient()

	var verifiedOrg VerifiedOrganization
	response, err := client.R().SetBody(signedPayload).SetResult(&verifiedOrg).Post("/identity")

	err = Utils().validateCreateResponse(response, err, "identity")
	return &verifiedOrg, err
}

// InvokeGet retrieve an organization
func (org *Organization) InvokeGet() (*VerifiedOrganization, error) {
	client := Utils().getDirectoryClient()

	nodeID := Utils().GenerateNodeID(org.Name)

	var verifiedOrg VerifiedOrganization
	response, err := client.R().SetResult(&verifiedOrg).Get("/orgs/" + nodeID)

	err = Utils().ValidateGetResponse(response, err, "org")
	return &verifiedOrg, err
}

// InvokeList retrieve a list of registered top-level organizations
func (org *Organization) InvokeList() (*[]ContractOrganization, error) {
	var orgs []ContractOrganization
	client := Utils().getNodeClient()
	instance, err := directory.NewDirectory(common.HexToAddress(Utils().getDirectoryAddress()), client)
	if err != nil {
		return &orgs, err
	}

	var rootNode [32]byte
	rootBytes, _ := hexutil.Decode(kaleido.RootNodeHash)
	copy(rootNode[:], rootBytes)

	count, err := instance.NodeChildrenCount(&bind.CallOpts{}, rootNode)
	if err != nil {
		return &orgs, err
	}
	countInt := count.Int64()
	var index int64
	fmt.Println("**********************************************************")
	fmt.Println("Number of orgs  =", count)
	for index = 0; index < countInt; index++ {
		var org ContractOrganization
		nodeID, _, err := instance.NodeChild(&bind.CallOpts{}, rootNode, uint8(index))
		if err != nil {
			return &orgs, err
		}
		owner, label, parent, _, _, _, _, err := instance.NodeDetails(&bind.CallOpts{}, nodeID)
		if err != nil {
			return &orgs, err
		}
		org.ID = "0x" + hex.EncodeToString(nodeID[:32])
		org.Name = label
		org.Owner = owner.String()
		org.ParentID = "0x" + hex.EncodeToString(parent[:32])
		orgs = append(orgs, org)
	}
	return &orgs, nil
}
