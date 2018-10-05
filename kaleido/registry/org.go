package registry

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"io/ioutil"
	"strings"

	jose "gopkg.in/square/go-jose.v2"
)

// Organization ...
type Organization struct {
	Consortium     string `json:"consortia_id,omitempty"`
	Environment    string `json:"environment_id,omitempty"`
	MemberID       string `json:"membership_id,omitempty"`
	Name           string
	Owner          string
	SigningKeyFile string
	CertPEMFile    string
}

// VerifiedOrganization ...
type VerifiedOrganization struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Owner    string `json:"owner,omitempty"`
	Proof    string `json:"proof,omitempty"`
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

	client := utils().getAPIClient()
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
	return noncePayload.Nonce, utils().validateCreateResponse(response, err, "nonce")
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
	x509Encoded := block.Bytes
	privateKey, err := x509.ParsePKCS8PrivateKey(x509Encoded)
	if err != nil {
		return nil, err
	}
	ecdsaKey := privateKey.(*ecdsa.PrivateKey)
	defer zeroKey(ecdsaKey)

	// read the provided proof
	proof, err := ioutil.ReadFile(org.CertPEMFile)
	if err != nil {
		return nil, err
	}

	// create a new signer using ECDSA (ES256) algorithm with the given private key
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.ES256, Key: ecdsaKey}, nil)
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
		"nonce":   nonce,
		"name":    org.Name,
		"proof":   string(proof),
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

// InvokeCreate registers a verified organization with the on-chain registry
// and stores the proof on-chain
func (org *Organization) InvokeCreate() (*VerifiedOrganization, error) {
	client := utils().getAPIClient()

	// sign payload
	signedPayload, err := org.createSignedRequestForRegistration()
	if err != nil {
		return nil, err
	}

	var verifiedOrg VerifiedOrganization
	response, err := client.R().SetBody(signedPayload).SetResult(&verifiedOrg).Post("/identity")

	err = utils().validateCreateResponse(response, err, "identity")
	return &verifiedOrg, err
}

// InvokeGet retrieve an organization
func (org *Organization) InvokeGet() (*VerifiedOrganization, error) {
	client := utils().getDirectoryClient()

	nodeID := utils().generateNodeID(org.Name)

	var verifiedOrg VerifiedOrganization
	response, err := client.R().SetResult(&verifiedOrg).Get("/orgs/" + nodeID)

	err = utils().validateGetResponse(response, err, "org")
	return &verifiedOrg, err
}

// InvokeList retrieve a list of registered top-level organizations
func (org *Organization) InvokeList() (*[]VerifiedOrganization, error) {
	type responseBodyType struct {
		Count int                    `json:"count,omitempty"`
		Orgs  []VerifiedOrganization `json:"orgs,omitempty"`
	}
	var responseBody responseBodyType
	client := utils().getDirectoryClient()
	response, err := client.R().SetResult(&responseBody).Get("/orgs")

	err = utils().validateGetResponse(response, err, "orgs")
	return &responseBody.Orgs, err
}
