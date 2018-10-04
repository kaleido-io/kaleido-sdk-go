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
	"github.com/kaleido-io/kaleido-sdk-go/kaleido/registry"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var orgsListCmd = &cobra.Command{
	Use:   "orgs",
	Short: "List the orgs",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		org := registry.Organization{}
		var verifiedOrgs *[]registry.VerifiedOrganization
		var err error
		if verifiedOrgs, err = org.InvokeList(); err != nil {
			cmd.SilenceUsage = true  // not a usage error at this point
			cmd.SilenceErrors = true // no need to display Error:, this still displays the error that is returned from RunE
			return err
		}
		jsonPrint(*verifiedOrgs)
		return nil
	},
}

var orgGetCmd = &cobra.Command{
	Use:   "org",
	Short: "Get the org details",
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
		if verified, err = org.InvokeGet(); err != nil {
			cmd.SilenceUsage = true  // not a usage error at this point
			cmd.SilenceErrors = true // no need to display Error:, this still displays the error that is returned from RunE
			return err
		}
		jsonPrint(*verified)
		return nil
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

var ownerAddress ethereumValue

func initCreateOrgCmd() {
	flags := orgCreateCmd.Flags()

	flags.StringP("memberid", "m", "", "Membership ID of the org")
	flags.StringP("proof", "p", "", "Path to identity certificate used when identifying organization on Kaleido")
	flags.StringP("key", "k", "", "Path to a key that should be used for signing the payload for registration")
	flags.VarP(&ethereumValue{}, "owner", "o", "Ethereum address for the owner of the organization")
	viper.BindPFlag("registry.create.org.memberid", flags.Lookup("memberid"))
	viper.BindPFlag("registry.create.org.proof", flags.Lookup("proof"))
	viper.BindPFlag("registry.create.org.key", flags.Lookup("key"))
	viper.BindPFlag("registry.create.org.owner", flags.Lookup("owner"))

	orgCreateCmd.MarkFlagRequired("memberid")
	orgCreateCmd.MarkFlagRequired("proof")
	orgCreateCmd.MarkFlagRequired("key")
	orgCreateCmd.MarkFlagRequired("owner")
}

func initGetOrgCmd() {
	flags := orgGetCmd.Flags()

	flags.StringP("memberid", "m", "", "Membership ID of the org")
	viper.BindPFlag("registry.get.org.memberid", flags.Lookup("memberid"))

	orgGetCmd.MarkFlagRequired("memberid")
}

func initListOrgCmd() {
	flags := orgsListCmd.Flags()

	flags.StringP("memberid", "m", "", "Membership ID of the org")
	viper.BindPFlag("registry.get.orgs.memberid", flags.Lookup("memberid"))

	orgsListCmd.MarkFlagRequired("memberid")
}

func init() {
	initCreateOrgCmd()
	initGetOrgCmd()
	initListOrgCmd()

	createCmd.AddCommand(orgCreateCmd)
	getCmd.AddCommand(orgGetCmd)
	getCmd.AddCommand(orgsListCmd)
}