// Copyright 2018 Kaleido, a ConsenSys business

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package registry

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	eth "github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	jose "gopkg.in/square/go-jose.v2"
)

// Organization ...
type Organization struct {
	Consortium     string `json:"consortia_id,omitempty"`
	Environment    string `json:"environment_id,omitempty"`
	MemberID       string `json:"membership_id,omitempty"`
	name           string
	owner          string
	signingKeyFile string
	certPEMFile    string
}

// VerifiedOrganization ...
type VerifiedOrganization struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Owner    string `json:"owner,omitempty"`
	Proof    string `json:"proof,omitempty"`
	ParentID string `json:"parent,omitempty"`
}

// NewOrganization creates a new organization from a command and its arguments
func NewOrganization(cmd *cobra.Command, args []string) *Organization {
	return &Organization{
		Consortium:     viper.GetString("consortium"),
		Environment:    viper.GetString("environment"),
		MemberID:       viper.GetString("memberid"),
		name:           args[0],
		owner:          viper.GetString("owner"),
		signingKeyFile: viper.GetString("key"),
		certPEMFile:    viper.GetString("proof"),
	}
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

	client := utils().getClient()
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
	pemEncodedBytes, err := ioutil.ReadFile(org.signingKeyFile)
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
	proof, err := ioutil.ReadFile(org.certPEMFile)
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
		"name":    org.name,
		"proof":   string(proof),
		"address": org.owner})

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

func (org *Organization) invokeCreate() (*VerifiedOrganization, error) {
	client := utils().getClient()

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

var orgsListCmd = &cobra.Command{
	Use:   "orgs",
	Short: "List the orgs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("orgs")
	},
}

var orgGetCmd = &cobra.Command{
	Use:   "orgs",
	Short: "Get the org details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Get org details")
	},
}

var orgCreateCmd = &cobra.Command{
	Use:     "org",
	Short:   "Create an on-chain organization",
	Example: "kld registry create org kaleido.com -c cid -e eid -p /path/to/proof/cert.pem -k /path/to/private/key -o 0xdEC89f82A6934DE1EA00CEa5A64233AdB898ACD8",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		org := NewOrganization(cmd, args)

		var verified *VerifiedOrganization
		var err error
		if verified, err = org.invokeCreate(); err != nil {
			cmd.SilenceUsage = true  // not a usage error at this point
			cmd.SilenceErrors = true // no need to display Error:, this still displays the error that is returned from RunE
			return err
		}
		jsonPrint(*verified)
		return nil
	},
}

type ethereumValue struct {
	address string
}

func (value *ethereumValue) String() string {
	return value.address
}

func (value *ethereumValue) Set(address string) error {
	if !eth.IsHexAddress(address) {
		return errors.New("owner must be a valid ethereum address")
	}
	value.address = address
	return nil
}

func (value *ethereumValue) Type() string {
	return "ethereum-address"
}

var ownerAddress ethereumValue

func init() {
	flags := orgCreateCmd.Flags()
	// —proof=/path/to/cert —path=/ --owner 0xasdfasdfasdfsadfdasdf
	flags.StringP("memberid", "m", "", "Membership ID of the org")
	flags.StringP("proof", "p", "", "Path to identity certificate used when identifying organization on Kaleido")
	flags.StringP("key", "k", "", "Path to a key that should be used for signing the payload for registration")
	flags.VarP(&ethereumValue{}, "owner", "o", "Ethereum address for the owner of the organization")
	viper.BindPFlag("memberid", flags.Lookup("memberid"))
	viper.BindPFlag("proof", flags.Lookup("proof"))
	viper.BindPFlag("key", flags.Lookup("key"))
	viper.BindPFlag("owner", flags.Lookup("owner"))

	orgCreateCmd.MarkFlagRequired("memberid")
	orgCreateCmd.MarkFlagRequired("proof")
	orgCreateCmd.MarkFlagRequired("key")
	orgCreateCmd.MarkFlagRequired("owner")

	createCmd.AddCommand(orgCreateCmd)
	getCmd.AddCommand(orgGetCmd)
	getCmd.AddCommand(orgsListCmd)
}
