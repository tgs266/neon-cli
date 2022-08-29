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
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tgs266/neon-cli/internal/neon"
	"github.com/tgs266/neon-cli/internal/parser"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		host, _ := cmd.Flags().GetString("host")
		filePaths := []string{}
		err := filepath.Walk(path,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if strings.HasSuffix(path, "yaml") || strings.HasSuffix(path, "yml") {
					filePaths = append(filePaths, path)
				}
				return nil
			})
		if err != nil {
			log.Println(err)
		}
		allManifests := []map[string]any{}
		for _, file := range filePaths {
			manifests := parser.ReadManifests(file)
			allManifests = append(allManifests, manifests...)
		}
		neon.Apply(host, allManifests)
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

	applyCmd.Flags().StringP("path", "p", "", "Path to manifests")
}
