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
