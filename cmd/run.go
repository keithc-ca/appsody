// Copyright © 2019 IBM Corporation and others.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

func newRunCmd(rootConfig *RootCommandConfig) *cobra.Command {
	config := &devCommonConfig{RootCommandConfig: rootConfig}
	// runCmd represents the run command
	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run your Appsody project in a containerized development environment.",
		Long: `Run the local Appsody environment, starting a container-based, continuous build environment for your project.
		
Run this command from the root directory of your Appsody project.`,
		Example: `  appsody run
  Runs your project in a containerized development environment.

  appsody run --interactive
  Runs your project in a containerized development environment, and attaches the standard input stream to the container. You can use the standard input stream to interact with processes inside the container.

  appsody run -p 3001:3000 --docker-options "--privileged" 
  Runs your project in a containerized development environment, binds the container port 3000 to the host port 3001, and passes the "--privileged" option to the "docker run" command as a flag.`,
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) > 0 {
				return errors.New("Unexpected argument. Use 'appsody [command] --help' for more information about a command")
			}
			rootConfig.Info.log("Running development environment...")
			return commonCmd(config, "run")

		},
	}

	addDevCommonFlags(runCmd, config)
	return runCmd
}
