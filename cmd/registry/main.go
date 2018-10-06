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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var registryCmd = &cobra.Command{
	Use:   "registry",
	Short: "Manage on-chain indentity registry",
}

func init() {
	// rootCmd adds registrCmd manually

	// create persistent flags for all sub-commands as this is not a complete command by itself
	registryCmd.PersistentFlags().StringP("service-id", "i", "", "Service ID (optional if config is setup properly)")
	registryCmd.PersistentFlags().StringP("consortium", "c", "", "Consortium ID")
	registryCmd.PersistentFlags().StringP("environment", "e", "", "Environment ID")

	registryCmd.MarkPersistentFlagRequired("consortium")
	registryCmd.MarkPersistentFlagRequired("environment")

	viper.BindPFlag("services.idregistry.id", registryCmd.PersistentFlags().Lookup("service-id"))
}

// NewRegistryCmd registry cmd
func NewRegistryCmd() *cobra.Command {
	return registryCmd
}
