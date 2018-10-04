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
package profile

import (
	"fmt"
	"strconv"
	"time"

	common "github.com/kaleido-io/kaleido-sdk-go/cmd/common"
	"github.com/kaleido-io/kaleido-sdk-go/kaleido/registry"
	"github.com/spf13/cobra"
)

var keysListCmd = &cobra.Command{
	Use:   "keys",
	Short: "List the keys for a given owner",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		profile := registry.Profile{}
		owner := cmd.Flags().Lookup("owner").Value.String()

		fmt.Println(owner)

		var properties *[]registry.Property
		var err error
		if properties, err = profile.GetProperties(owner); err != nil {
			return err
		}
		common.PrintJSON(properties)
		return nil
	},
}

var keyGetCmd = &cobra.Command{
	Use:   "key",
	Short: "Get value(s) for a key for a given owner",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		revision, _ := cmd.Flags().GetString("revision")
		profile := registry.Profile{}
		owner := cmd.Flags().Lookup("owner").Value.String()

		cmd.SilenceErrors = true
		cmd.SilenceUsage = true
		key := args[0]
		if revision == "" {
			property, err := profile.GetProperty(owner, key)
			if err != nil {
				return err
			}
			common.PrintJSON(property)
		} else if revision == "all" {
			properties, err := profile.GetPropertyAllVersions(owner, key)
			if err != nil {
				return err
			}
			common.PrintJSON(properties)
		} else {
			revisionIndex, err := strconv.ParseInt(revision, 10, 64)
			if err != nil {
				return err
			}

			property, err := profile.GetPropertyByRevision(owner, key, revisionIndex)
			if err != nil {
				return err
			}
			common.PrintJSON(property)
		}
		return nil
	},
}

var keySetCmd = &cobra.Command{
	Use:   "key",
	Short: "Set the key to a particular value",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		owner := cmd.Flags().Lookup("owner").Value.String()
		var keystorePath string
		var err error
		if keystorePath, err = cmd.Flags().GetString("keystore"); err != nil {
			return err
		}
		profile := registry.Profile{
			KeyStorePath: keystorePath,
			Signer:       owner,
		}

		revision, err := cmd.Flags().GetString("revision")
		if err != nil {
			return err
		}

		value, err := cmd.Flags().GetString("value")
		if err != nil {
			return err
		}

		if err = profile.SetProperty(args[0], value, revision); err != nil {
			cmd.SilenceErrors = true
			cmd.SilenceUsage = true
			return err
		}

		return nil
	},
}

func initSetKeyCmd() {
	flags := keySetCmd.Flags()

	now := time.Now().Unix()
	flags.StringP("value", "v", "", "Value for the key")
	flags.VarP(&common.EthereumAddress{}, "owner", "o", "Account of the profile owner")
	flags.StringP("keystore", "k", "", "Keystore path so accounts can be used to sign tx")
	flags.StringP("revision", "n", strconv.FormatInt(now, 10), "Revision for the key")

	keySetCmd.MarkFlagRequired("value")
	keySetCmd.MarkFlagRequired("signer")
	keySetCmd.MarkFlagRequired("keystore")
}

func initGetKeyCmd() {
	flags := keyGetCmd.Flags()

	flags.VarP(&common.EthereumAddress{}, "owner", "o", "Ethereum account of the owner of the profile")
	flags.StringP("revision", "n", "", "Revision for the key")

	keyGetCmd.MarkFlagRequired("owner")
}

func initListKeysCmd() {
	flags := keysListCmd.Flags()

	flags.VarP(&common.EthereumAddress{}, "owner", "o", "Ethereum account of the owner of the profile")

	keysListCmd.MarkFlagRequired("owner")
}

func init() {
	initSetKeyCmd()
	initGetKeyCmd()
	initListKeysCmd()

	setCmd.AddCommand(keySetCmd)
	getCmd.AddCommand(keyGetCmd)
	getCmd.AddCommand(keysListCmd)
}
