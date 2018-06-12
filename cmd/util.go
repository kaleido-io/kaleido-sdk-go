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

	kld "github.com/kaleido-io/kaleido-sdk-go/kaleido"
	"github.com/spf13/viper"
)

var getNewClient = func() kld.KaleidoClient {
	return kld.NewClient(viper.GetString("api.url"), viper.GetString("api.key"))
}

func validateName() {
	if name == "" {
		fmt.Println("Missing required parameter: name")
		os.Exit(1)
	}

	return
}

func validateProvider() {
	if provider == "" {
		fmt.Println("Missing required parameter: provider (quorum or geth)")
		os.Exit(1)
	}

	return
}

func validateConsensus() {
	if consensus == "" {
		fmt.Println("Missing required parameter: consensus (raft or ibft for quorum, or poa for geth)")
		os.Exit(1)
	}

	return
}

func validateConsortiumId(resourceName string, isCreate ...bool) {
	if consortiumId == "" {
		if len(isCreate) == 0 || isCreate[0] {
			fmt.Printf("Missing required parameter: --consortiumId for the consortium to add the new %s to\n", resourceName)
		} else {
			fmt.Printf("Missing required parameter: --consortiumId for the consortium to delete the %s from\n", resourceName)
		}

		os.Exit(1)
	}
}

func validateAppCredsId(resourceName string, isGet ...bool) {
	if appCredsId == "" {
		if len(isGet) == 0 || isGet[0] {
			fmt.Printf("Missing required parameter: --appcredsId to get the appilcation credentials %s to\n", resourceName)
		} else {
			fmt.Printf("Missing required parameter: --appcredsId to delete the appilcation credentials %s from\n", resourceName)
		}

		os.Exit(1)
	}
}

func validateEnvironmentId(resourceName string, isCreate ...bool) {
	if environmentId == "" {
		if len(isCreate) == 0 || isCreate[0] {
			fmt.Printf("Missing required parameter: --environmentId for the environment to add the new %s to\n", resourceName)
		} else {
			fmt.Printf("Missing required parameter: --environmentId for the environment to delete the %s from\n", resourceName)
		}

		os.Exit(1)
	}
}

func validateMembershipId(resourceName string, isCreate ...bool) {
	if membershipId == "" {
		if len(isCreate) == 0 || isCreate[0] {
			fmt.Printf("Missing required parameter: --membershipId for the membership to add the new %s to\n", resourceName)
		} else {
			fmt.Printf("Missing required parameter: --membershipId for the membership to delete the %s from\n", resourceName)
		}

		os.Exit(1)
	}
}

func validateDeleteId(resourceName string) {
	if deleteId == "" {
		fmt.Printf("Missing required parameter: --id for the %s to delete\n", resourceName)
		os.Exit(1)
	}

	return
}
