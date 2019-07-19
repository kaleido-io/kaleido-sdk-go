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
	"errors"
	"fmt"
	"strconv"
	"strings"

	kaleido "github.com/kaleido-io/kaleido-sdk-go/kaleido/registry"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Manage on-chain identity profile",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		serviceID := viper.GetString("services.idregistry.id")
		if serviceID == "" {
			// not specified via environment, let's check our flag
			var err error
			if serviceID, err = cmd.Flags().GetString("service-id"); err != nil || serviceID == "" {
				// yeah, we need it so error out
				fmt.Println(err)
				return errors.New("missing service id. have you setup the config file (~/.kld.yaml) or did you specify --service-id")
			}
		}
		// at this point, we should have a serviceID, setup viper so other components can access it without access to cmd
		viper.Set("service.id", serviceID)

		// Check for compatibility between CLI and deployed smart contract versions
		var service *kaleido.ServiceDefinitionType
		var err error
		if service, err = kaleido.Utils().GetServiceDefinition(); err != nil {
			return err
		}
		clientNetMgr := kaleido.Utils().GetNetworkManagerClient()
		targetURL := "/consortia/" + service.Consortium + "/environments/" + service.Environment + "/services/" + serviceID + "/status"
		type urlsJSON struct {
			HTTP string `json:"http,omitempty"`
		}
		type responseBody struct {
			Urls            urlsJSON `json:"urls,omitempty"`
			Status          string   `json:"status,omitempty"`
			Release         string   `json:"release,omitempty"`
			ContractVersion string   `json:"contract_version,omitempty"`
		}
		var servicePayload responseBody
		serviceResponse, err := clientNetMgr.R().
			SetHeader("Content-Type", "application/json").
			SetResult(&servicePayload).
			Get(targetURL)

		err = kaleido.Utils().ValidateGetResponse(serviceResponse, err, "service")
		if err != nil {
			return err
		}

		contractVersion := servicePayload.ContractVersion
		if len(contractVersion) == 0 {
			fmt.Println("Contracts version not returned so will assume version = 1.x.x")
		} else {
			fmt.Println("Detected version " + contractVersion + " of the deployed smart contracts.")
		}
		tokens := strings.Split(contractVersion, ".")
		majorVersion, _ := strconv.Atoi(tokens[0])
		if majorVersion != 2 {
			return errors.New("This version of kaleido-sdk-go CLI is not supported for your deployed smart contract versions. Please upgrade to a CLI compatible with " + contractVersion)
		}

		fmt.Println("This version of the CLI is compatible.")

		return nil
	}}

func init() {
	// rootCmd adds registrCmd manually

	// create persistent flags for all sub-commands as this is not a complete command by itself
	profileCmd.PersistentFlags().StringP("service-id", "i", "", "Service ID (optional if config is setup properly)")
	viper.BindPFlag("profile", profileCmd.PersistentFlags().Lookup("profile"))
}

// NewProfileCmd registry cmd
func NewProfileCmd() *cobra.Command {
	return profileCmd
}
