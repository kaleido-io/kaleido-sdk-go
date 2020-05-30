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

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/kaleido-io/kaleido-sdk-go/cmd/profile"

	"github.com/kaleido-io/kaleido-sdk-go/cmd/registry"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// for create command
var name string
var desc string
var provider string
var consensus string
var serviceType string
var email string
var region string
var cloud string
var multiRegion bool
var blockPeriod int
var configType string
var detailsFile string
var size string
var kmsID string
var opsmetricID string
var backupID string
var networkingID string
var nodeConfigID string
var bafID string

// use for both create, list, get and delete commands
var consortiumID string
var czoneID string
var membershipID string
var environmentID string
var ezoneID string
var nodeID string
var serviceID string
var appCredsID string
var service string
var invitationID string
var configID string

// for delete command
var deleteID string

var rootCmd = &cobra.Command{
	Use:   "kld",
	Short: "Command Line Tool for Kaleido resources management",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("API URL: %s\n", viper.Get("api.url"))
	},
}

var cfgFile string

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// all environment variables for the "kld" command will have the "KLD" prefix
	// e.g "KLD_API_URL"
	viper.SetEnvPrefix("kld")
	// allows code to access env variables by name only, without the prefix
	viper.AutomaticEnv()
	// allows using "." to access env variables with "_"
	// e.g viper.Get('api.url') for value of "KLD_API_URL"
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	rootCmd.PersistentFlags().Int64("verbose", 0, "Verbosity level of output (0 or 1)")

	rootCmd.PersistentFlags().String("api-url", "", "Kaleido API URL (optional)")
	viper.BindPFlag("api.url", rootCmd.PersistentFlags().Lookup("api-url"))

	rootCmd.PersistentFlags().String("api-key", "", "Kaleido API KEY (optional)")
	viper.BindPFlag("api.key", rootCmd.PersistentFlags().Lookup("api-key"))

	// config files capture defaults that can be overwritten by env variables and flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file that captures re-usable settings such as API URl, API Key, etc. (default is $HOME/.kld.yaml)")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	rootCmd.AddCommand(newCreateCmd())
	rootCmd.AddCommand(newDeleteCmd())
	rootCmd.AddCommand(newListCmd())
	rootCmd.AddCommand(newGetCmd())

	// add registry command
	rootCmd.AddCommand(registry.NewRegistryCmd())
	rootCmd.AddCommand(profile.NewProfileCmd())
}

func initConfig() {
	verbose := 0
	if verbose > 1 {
		fmt.Println("initializing config")
	}
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".kld" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".kld")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("\nCan't read config: %v, will rely on environment variables for required configurations\n", err)
	}

	viper.SetDefault("api.debug", false)
}
