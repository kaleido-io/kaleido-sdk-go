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

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List various resources this user account owns: consortium, membership, environment, node, appcreds",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list command")
	},
}

func newListCmd() *cobra.Command {
	listCmd.AddCommand(newConsortiumListCmd())
	listCmd.AddCommand(newEnvironmentListCmd())
	listCmd.AddCommand(newNodeListCmd())
	listCmd.AddCommand(newServiceListCmd())
	listCmd.AddCommand(newMembershipListCmd())
	listCmd.AddCommand(newAppCredsListCmd())
	listCmd.AddCommand(newInvitationListCmd())

	return listCmd
}
