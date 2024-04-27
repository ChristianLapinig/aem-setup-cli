/*
Copyright © 2024 Christian Lapinig <lapinig.a.christian@gmail.com>
*/
package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/ChristianLapinig/aem-starter/lib"
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
		quickstartPath, _ := cmd.Flags().GetString("quickstart")

		authorPath := filepath.Join(setupPath, "author")
		publishPath := filepath.Join(setupPath, "publish")

		fmt.Println("Creating author and publish folders in", setupPath)
		if err := lib.CreateFolder(authorPath); err != nil {
			panic(err)
		}
		fmt.Println("Successfully created author folder")

		if err := lib.CreateFolder(publishPath); err != nil {
			panic(err)
		}
		fmt.Println("Successfully created publish folder")

		// TODO: It should be specified where to mount the given repository
		// should it be in the author or publish folder? or both?

		fmt.Println("Copying quickstart JAR to respective folders")
		if quickstartPath != "" {
			fmt.Println(authorPath + "/aem-author-p4502.jar")
			fmt.Println(publishPath + "/aem-publish-p4503.jar")
		}
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
	setupCmd.Flags().StringP("path", "p", ".", "Path where AEM should be setup. Default is the current directory you are in.")
	setupCmd.Flags().StringP("mount", "m", "", "Path to an existing AEM repository that will be mounted to your instances.")
	setupCmd.Flags().StringP("quickstart", "q", "", "Path to an existing AEM quickstart JAR. For example, the latest AEM Cloud SDK quickstart JAR.")
}
