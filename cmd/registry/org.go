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
	"errors"
	"fmt"

	eth "github.com/ethereum/go-ethereum/common"
	"github.com/kaleido-io/kaleido-sdk-go/kaleido/registry"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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
		org := registry.NewOrganization(cmd, args)

		var verified *registry.VerifiedOrganization
		var err error
		if verified, err = org.InvokeCreate(); err != nil {
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
