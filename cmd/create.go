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

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create various resources: consortium, membership, environment, node, appcreds, invitation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create command")
	},
}

func newCreateCmd() *cobra.Command {
	createCmd.AddCommand(newConsortiumCreateCmd())
	createCmd.AddCommand(newMembershipCreateCmd())
	createCmd.AddCommand(newEnvironmentCreateCmd())
	createCmd.AddCommand(newNodeCreateCmd())
	createCmd.AddCommand(newServiceCreateCmd())
	createCmd.AddCommand(newAppCredsCreateCmd())
	createCmd.AddCommand(newInvitationCreateCmd())

	return createCmd
}
