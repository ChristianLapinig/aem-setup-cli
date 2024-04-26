/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Creates the necessary files and folders for a local AEM environment.",
	Long:  `'setup' creates the necessary files and folders to run a local AEM environment.`,
	Run: func(cmd *cobra.Command, args []string) {
		setupPath, _ := cmd.Flags().GetString("path")
		// repoPath, _ := cmd.Flags().GetString("mount")
		// quickstartPath, _ := cmd.Flags().GetString("quickstart")

		authorPath := filepath.Join(setupPath, "author")
		publishPath := filepath.Join(setupPath, "publish")

		fmt.Println("Creating author and publish folders in", setupPath)
		if err := os.Mkdir(authorPath, 0755); err != nil {
			panic(err)
		}
		fmt.Println("Successfully created author folder")

		if err := os.Mkdir(publishPath, 0755); err != nil {
			panic(err)
		}
		fmt.Println("Successfully created publish folder")
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
	setupCmd.Flags().StringP("path", "p", ".", "Path where AEM should be setup. Default is the current directory you are in.")
	setupCmd.Flags().StringP("mount", "m", "", "Path to an existing AEM repository that will be mounted to your instances.")
	setupCmd.Flags().StringP("quickstart", "q", "", "Path to an existing AEM quickstart JAR. For example, the latest AEM Cloud SDK quickstart JAR.")
}
