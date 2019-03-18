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

var accountGetCmd = &cobra.Command{
	Use:   "account",
	Short: "Get an account's details",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		parent := cmd.Flags().Lookup("parent").Value.String()
		if parent[:2] != "0x" && parent[:1] != "/" {
			return errors.New("flag 'parent' value must start with either a '0x' or a '/'")
		}
		acct := &registry.Account{
			Parent: parent,
			Name:   args[0],
		}
		var err error
		if err = acct.InvokeGet(); err != nil {
			cmd.SilenceUsage = true  // not a usage error at this point
			cmd.SilenceErrors = true // no need to display Error:, this still displays the error that is returned from RunE
			return err
		}
		return nil
	},
}

var accountCreateCmd = &cobra.Command{
	Use:   "account",
	Short: "Add an account address to an org or group",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		parent := cmd.Flags().Lookup("parent").Value.String()
		if parent[:2] != "0x" && parent[:1] != "/" {
			return errors.New("flag 'parent' value must start with either a '0x' or a '/'")
		}
		address := cmd.Flags().Lookup("address").Value.String()
		if address[:2] != "0x" {
			return errors.New("flag 'address' value must start with '0x'")
		}

		acct := &registry.Account{
			Parent: parent,
			Name:   args[0],
			Value:  address,
		}

		var keystorePath string
		var signer string

		keystorePath = cmd.Flags().Lookup("keystore").Value.String()
		signer = cmd.Flags().Lookup("signer").Value.String()

		var err error
		if err = acct.InvokeCreate(keystorePath, signer); err != nil {
			cmd.SilenceUsage = true  // not a usage error at this point
			cmd.SilenceErrors = true // no need to display Error:, this still displays the error that is returned from RunE
			return err
		}
		return nil
	},
}

func initCreateAccountCmd() {
	flags := accountCreateCmd.Flags()

	flags.StringP("parent", "p", "", "Name path to the parent org or group")
	flags.VarP(&common.EthereumAddress{}, "signer", "s", "Account owned by parent (used to sign tx)")
	flags.StringP("keystore", "k", "", "Keystore directory path so account can be used to sign tx")
	flags.VarP(&common.EthereumAddress{}, "address", "a", "Ethereum account address to store")

	accountCreateCmd.MarkFlagRequired("parent")
	accountCreateCmd.MarkFlagRequired("signer")
	accountCreateCmd.MarkFlagRequired("keystore")
	accountCreateCmd.MarkFlagRequired("address")
}

func initGetAccountCmd() {
	flags := accountGetCmd.Flags()

	flags.StringP("parent", "p", "", "Name path to the parent org or group")

	accountGetCmd.MarkFlagRequired("parent")
}

func init() {
	initCreateAccountCmd()
	initGetAccountCmd()

	createCmd.AddCommand(accountCreateCmd)
	getCmd.AddCommand(accountGetCmd)
}
