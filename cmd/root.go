/*
Copyright © 2024 Christian Lapinig <lapinig.a.christian@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aem-starter",
	Short: "CLI tool to help setup AEM locally",
	Long:  `CLI tool that makes it easy to setup your environment to run AEM locally.`,
	Run: func(cmd *cobra.Command, args []string) {
		setupPath, _ := cmd.Flags().GetString("path")
		repoPath, _ := cmd.Flags().GetString("mount")
		fmt.Println("Setup path: ", setupPath)

		if repoPath != "" {
			fmt.Println("Repository path: ", repoPath)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("path", "p", ".", "Path where AEM should be setup. Default is the current directory you are in.")
	rootCmd.Flags().StringP("mount", "m", "", "Path to an existing AEM repository that will be mounted to the author and publish instances.")
}
