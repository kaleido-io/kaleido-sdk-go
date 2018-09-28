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
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var groupsListCmd = &cobra.Command{
	Use:   "groups",
	Short: "List the groups within an org",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list groups")
	},
}

var groupGetCmd = &cobra.Command{
	Use:   "group",
	Short: "Get the group details identified by a path",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Get group details")
	},
}

var groupCreateCmd = &cobra.Command{
	Use:   "group",
	Short: "Create a group at the given path (for an org)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Missing required argument(s)")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create group")
	},
}

func init() {
	createCmd.AddCommand(groupCreateCmd)
	getCmd.AddCommand(groupGetCmd)
	getCmd.AddCommand(groupsListCmd)
}
