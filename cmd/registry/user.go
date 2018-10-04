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
	"github.com/kaleido-io/kaleido-sdk-go/common"
	"github.com/kaleido-io/kaleido-sdk-go/kaleido/registry"
	"github.com/spf13/cobra"
)

var usersListCmd = &cobra.Command{
	Use:   "users",
	Short: "List the users within an org",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		user := &registry.User{
			Parent: cmd.Flags().Lookup("parent").Value.String(),
		}

		var users *[]registry.User
		if users, err = user.InvokeList(); err != nil {
			cmd.SilenceUsage = true  // not a usage error at this point
			cmd.SilenceErrors = true // no need to display Error:, this still displays the error that is returned from RunE
			return err
		}
		common.PrintJSON(users)
		return nil
	},
}

var userGetCmd = &cobra.Command{
	Use:   "user",
	Short: "Get the user details identified by a path",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		user := &registry.User{
			Parent: cmd.Flags().Lookup("parent").Value.String(),
			Email:  args[0],
		}

		var err error
		if user, err = user.InvokeGet(); err != nil {
			cmd.SilenceUsage = true  // not a usage error at this point
			cmd.SilenceErrors = true // no need to display Error:, this still displays the error that is returned from RunE
			return err
		}
		common.PrintJSON(user)
		return nil
	},
}

var userCreateCmd = &cobra.Command{
	Use:   "user",
	Short: "Create a user at the given path (for an org)",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var user *registry.User
		user = &registry.User{
			Email:       args[0],
			Consortium:  cmd.Flags().Lookup("consortium").Value.String(),
			Environment: cmd.Flags().Lookup("environment").Value.String(),
			MemberID:    cmd.Flags().Lookup("memberid").Value.String(),
			Parent:      cmd.Flags().Lookup("parent").Value.String(),
			Owner:       cmd.Flags().Lookup("owner").Value.String(),
		}

		var keystorePath string
		var signer string

		keystorePath = cmd.Flags().Lookup("keystore").Value.String()
		signer = cmd.Flags().Lookup("signer").Value.String()

		var err error
		if err = user.InvokeCreate(keystorePath, signer); err != nil {
			cmd.SilenceUsage = true  // not a usage error at this point
			cmd.SilenceErrors = true // no need to display Error:, this still displays the error that is returned from RunE
			return err
		}
		return nil
	},
}

func initCreateUserCmd() {
	flags := userCreateCmd.Flags()

	flags.StringP("memberid", "m", "", "Membership ID of the org")
	flags.VarP(&common.EthereumAddress{}, "owner", "o", "Ethereum account that the user owns")
	flags.StringP("keystore", "k", "", "Keystore path so accounts can be used to sign tx")
	flags.VarP(&common.EthereumAddress{}, "signer", "s", "Account to use to sign tx")
	flags.StringP("parent", "p", "", "Path to the parent org or group")

	userCreateCmd.MarkFlagRequired("memberid")
	userCreateCmd.MarkFlagRequired("account")
	userCreateCmd.MarkFlagRequired("key")
	userCreateCmd.MarkFlagRequired("parent")
}

func initGetUserCmd() {
	flags := userGetCmd.Flags()

	flags.StringP("memberid", "m", "", "Membership ID of the org")
	flags.StringP("parent", "p", "", "Path to the parent org or group")

	userCreateCmd.MarkFlagRequired("memberid")
	userCreateCmd.MarkFlagRequired("parent")
}

func initListUserCmd() {
	flags := usersListCmd.Flags()

	flags.StringP("memberid", "m", "", "Membership ID of the org")
	flags.StringP("parent", "p", "", "Path to the parent org or group")

	userCreateCmd.MarkFlagRequired("memberid")
	userCreateCmd.MarkFlagRequired("parent")
}

func init() {
	initCreateUserCmd()
	initGetUserCmd()
	initListUserCmd()

	createCmd.AddCommand(userCreateCmd)
	getCmd.AddCommand(userGetCmd)
	getCmd.AddCommand(usersListCmd)
}
