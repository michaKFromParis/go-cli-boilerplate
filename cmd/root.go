// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/michaKFromParis/go-cli-boilerplate/logger"
	"github.com/spf13/cobra"
)

var verbose = false
var veryVerbose = false

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	PersistentPreRun: PreRunAllCommands,
	Use:              filepath.Base(os.Args[0]),
	Long: filepath.Base(os.Args[0]) + `
	
	A boilerplate for go command line applications
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// PreRunAllCommands is called by cobra for every commands just before executing a command
func PreRunAllCommands(cmd *cobra.Command, args []string) {
	if veryVerbose {
		logger.Init(log.TraceLevel)
	} else if verbose {
		logger.Init(log.DebugLevel)
	} else {
		logger.Init(log.InfoLevel)
	}
}

// SetVersion passes the current version to the cli
func SetVersion(version string, commit string, date string) {
	rootCmd.Version = version
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().SortFlags = false
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "set log verbose level to debug")
	rootCmd.PersistentFlags().BoolVarP(&veryVerbose, "v", "", false, "set log verbose level to trace")
}
