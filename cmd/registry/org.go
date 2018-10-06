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

	"github.com/kaleido-io/kaleido-sdk-go/common"
	"github.com/kaleido-io/kaleido-sdk-go/kaleido/registry"
	"github.com/spf13/cobra"
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
		common.PrintJSON(*verifiedOrgs)
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

		// arg[0] is the name and must start with a 0x or a /
		name := args[0]
		if name[:2] != "0x" && name[:1] != "/" {
			return errors.New("name of an org must being with a 0x or must be specified as a path beginning with /")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		org := &registry.Organization{
			Consortium:  cmd.Flags().Lookup("consortium").Value.String(),
			Environment: cmd.Flags().Lookup("environment").Value.String(),
			MemberID:    cmd.Flags().Lookup("memberid").Value.String(),
			Name:        args[0],
		}
		var verified *registry.VerifiedOrganization
		var err error
		if verified, err = org.InvokeGet(); err != nil {
			cmd.SilenceUsage = true  // not a usage error at this point
			cmd.SilenceErrors = true // no need to display Error:, this still displays the error that is returned from RunE
			return err
		}
		common.PrintJSON(*verified)
		return nil
	},
}

var orgCreateCmd = &cobra.Command{
	Use:   "org",
	Short: "Create an on-chain organization",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		org := &registry.Organization{
			Consortium:     cmd.Flags().Lookup("consortium").Value.String(),
			Environment:    cmd.Flags().Lookup("environment").Value.String(),
			MemberID:       cmd.Flags().Lookup("memberid").Value.String(),
			Name:           args[0],
			Owner:          cmd.Flags().Lookup("owner").Value.String(),
			SigningKeyFile: cmd.Flags().Lookup("pkcs8-key").Value.String(),
			CertPEMFile:    cmd.Flags().Lookup("proof").Value.String(),
		}

		var verified *registry.VerifiedOrganization
		var err error
		if verified, err = org.InvokeCreate(); err != nil {
			cmd.SilenceUsage = true  // not a usage error at this point
			cmd.SilenceErrors = true // no need to display Error:, this still displays the error that is returned from RunE
			return err
		}
		common.PrintJSON(*verified)
		return nil
	},
}

func initCreateOrgCmd() {
	flags := orgCreateCmd.Flags()

	flags.StringP("memberid", "m", "", "Membership ID of the org")
	flags.StringP("proof", "p", "", "Path to identity certificate used when identifying organization on Kaleido")
	flags.StringP("pkcs8-key", "k", "", "Path to a key that should be used for signing the payload for registration")
	flags.VarP(&common.EthereumAddress{}, "owner", "o", "Ethereum address for the owner of the organization")

	orgCreateCmd.MarkFlagRequired("memberid")
	orgCreateCmd.MarkFlagRequired("proof")
	orgCreateCmd.MarkFlagRequired("pkcs8-key")
	orgCreateCmd.MarkFlagRequired("owner")
}

func initGetOrgCmd() {
	flags := orgGetCmd.Flags()

	flags.StringP("memberid", "m", "", "Membership ID of the org")

	orgGetCmd.MarkFlagRequired("memberid")
}

func initListOrgCmd() {
	flags := orgsListCmd.Flags()

	flags.StringP("memberid", "m", "", "Membership ID of the org")

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
