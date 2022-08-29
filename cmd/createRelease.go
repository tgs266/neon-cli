/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tgs266/neon-cli/internal/neon"
	"github.com/tgs266/neon-cli/internal/parser"
)

// createReleaseCmd represents the createRelease command
var createReleaseCmd = &cobra.Command{
	Use:   "create-release",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		manifestPath, _ := cmd.Flags().GetString("manifest")
		host, _ := cmd.Flags().GetString("host")
		manifests := parser.ReadManifests(manifestPath)
		neon.Apply(host, manifests)
	},
}

func init() {
	rootCmd.AddCommand(createReleaseCmd)
	createReleaseCmd.Flags().StringP("manifest", "m", "", "Release manifest file")
}