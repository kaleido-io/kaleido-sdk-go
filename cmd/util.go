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
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gopkg.in/resty.v1"
)

func getNewClient() kld.KaleidoClient {
	if viper.GetString("api.url") == "" || viper.GetString("api.key") == "" {
		fmt.Println("Missing api-url or api-key. Have you setup your config (~/.kld.yaml), or env variables, or specified url and key via a flag?")
		os.Exit(1)
	}
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

func validateServiceType() {
	if serviceType == "" {
		fmt.Println("Missing required parameter: service")
		os.Exit(1)
	}

	return
}

func validateEmail() {
	if email == "" {
		fmt.Println("Missing required parameter: email")
		os.Exit(1)
	}

	return
}

func validateConsortiumID(resourceName string, isCreate ...bool) {
	if consortiumID == "" {
		if len(isCreate) == 0 || isCreate[0] {
			fmt.Printf("Missing required parameter: --consortiumID for the consortium to add the new %s to\n", resourceName)
		} else {
			fmt.Printf("Missing required parameter: --consortiumID for the consortium to delete the %s from\n", resourceName)
		}

		os.Exit(1)
	}
}

func validateAppCredsID(resourceName string, isGet ...bool) {
	if appCredsID == "" {
		if len(isGet) == 0 || isGet[0] {
			fmt.Printf("Missing required parameter: --appcredsID to get the appilcation credentials %s to\n", resourceName)
		} else {
			fmt.Printf("Missing required parameter: --appcredsID to delete the appilcation credentials %s from\n", resourceName)
		}

		os.Exit(1)
	}
}

func validateEnvironmentID(resourceName string, isCreate ...bool) {
	if environmentID == "" {
		if len(isCreate) == 0 || isCreate[0] {
			fmt.Printf("Missing required parameter: --environmentID for the environment to add the new %s to\n", resourceName)
		} else {
			fmt.Printf("Missing required parameter: --environmentID for the environment to delete the %s from\n", resourceName)
		}

		os.Exit(1)
	}
}

func validateMembershipID(resourceName string, isCreate ...bool) {
	if membershipID == "" {
		if len(isCreate) == 0 || isCreate[0] {
			fmt.Printf("Missing required parameter: --membershipID for the membership to add the new %s to\n", resourceName)
		} else {
			fmt.Printf("Missing required parameter: --membershipID for the membership to delete the %s from\n", resourceName)
		}

		os.Exit(1)
	}
}

func validateInvitationID(resourceName string, isCreate ...bool) {
	if invitationID == "" {
		if len(isCreate) == 0 || isCreate[0] {
			fmt.Printf("Missing required parameter: --invitationID for the invitation to add the new %s to\n", resourceName)
		} else {
			fmt.Printf("Missing required parameter: --invitationID for the invitation to delete the %s from\n", resourceName)
		}

		os.Exit(1)
	}
}

func validateDeleteID(resourceName string) {
	if deleteID == "" {
		fmt.Printf("Missing required parameter: --id for the %s to delete\n", resourceName)
		os.Exit(1)
	}

	return
}

func validateGetResponse(res *resty.Response, err error, resourceName string) {
	if res.StatusCode() != 200 {
		fmt.Printf("Could not retrieve %s. Status code: %d.", resourceName, res.StatusCode())
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\n%+v\n", res)
	}
}

func validateCreationResponse(res *resty.Response, err error, resourceName string) {
	if res != nil && res.StatusCode() != 201 {
		fmt.Printf("Could not create %s. Status code: %d.", resourceName, res.StatusCode())
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\n%+v\n", res)
	}
}

func validateDeletionResponse(res *resty.Response, err error, resourceName string) {
	if res.StatusCode() != 202 && res.StatusCode() != 204 {
		fmt.Printf("%s deletion failed. Status code: %d\n", resourceName, res.StatusCode())
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\n%+v\n", res)
	}
}

func printGetResponse(res *resty.Response, err error, resourceName string) error {
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)
	if res.StatusCode() != 200 {
		return errors.Errorf("Could not retrieve %s. Status code: %d\n", resourceName, res.StatusCode())
	}

	return nil
}

func printCreationResponse(res *resty.Response, err error, resourceName string) error {
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)
	if res.StatusCode() != 201 {
		return errors.Errorf("Could not create %s. Status code: %d\n", resourceName, res.StatusCode())
	}
	return nil
}

func printDeletionResponse(res *resty.Response, err error, resourceName string) error {
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)
	if res.StatusCode() != 202 && res.StatusCode() != 204 {
		return errors.Errorf("%s deletion failed. Status code: %d\n", resourceName, res.StatusCode())
	}
	return nil
}
