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
	"fmt"

	"github.com/spf13/cobra"
)

var usersListCmd = &cobra.Command{
	Use:   "users",
	Short: "List the users within an org",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list users")
	},
}

var userGetCmd = &cobra.Command{
	Use:   "user",
	Short: "Get the user details identified by a path",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Get user details")
	},
}

var userCreateCmd = &cobra.Command{
	Use:   "user",
	Short: "Create a user at the given path (for an org)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create user")
	},
}

func init() {
	createCmd.AddCommand(userCreateCmd)
	getCmd.AddCommand(userGetCmd)
	getCmd.AddCommand(usersListCmd)
}
