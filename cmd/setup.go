/*
Copyright © 2024 Christian Lapinig <lapinig.a.christian@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
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
		// TODO: Copy license.properties file. Needs to be an argument
		if len(args) < 1 {
			fmt.Println("license.properties file is required. Exiting...")
			os.Exit(1)
		}

		licenseProperties := args[0]
		setupPath, _ := cmd.Flags().GetString("path")
		quickstartPath, _ := cmd.Flags().GetString("quickstart")
		// repoPath, _ := cmd.Flags().GetString("mount")

		authorPath := filepath.Join(setupPath, "author")
		publishPath := filepath.Join(setupPath, "publish")

		fmt.Println("Creating author and publish folders in", setupPath)
		if err := lib.CreateFolder(authorPath); err != nil {
			panic(err)
		}

		if err := lib.CreateFolder(publishPath); err != nil {
			panic(err)
		}
		fmt.Println("Successfully created author and publish folders")

		fmt.Println("Copying license.properties file to respective instances")
		if err := lib.CopyFile(licenseProperties, filepath.Join(authorPath, "license.properties")); err != nil {
			panic(err)
		}

		if err := lib.CopyFile(licenseProperties, filepath.Join(publishPath, "license.properties")); err != nil {
			panic(err)
		}
		fmt.Println("Successfully copied license.properties to author and publish folders")

		// TODO: It should be specified where to mount the given repository
		// should it be in the author or publish folder? or both?

		fmt.Println("Copying quickstart JAR to respective folders")
		if err := lib.CopyFile(quickstartPath, filepath.Join(authorPath, "aem-author-p4502.jar")); err != nil {
			panic(err)
		}

		if err := lib.CopyFile(quickstartPath, filepath.Join(publishPath, "aem-publish-p4503.jar")); err != nil {
			panic(err)
		}
		fmt.Println("Successfully copied the quickstart jar to author and publish folder")
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
	setupCmd.Flags().StringP("path", "p", ".", "Path where AEM should be setup. Default is the current directory you are in.")
	setupCmd.Flags().StringP("quickstart", "q", "cq-quickstart-6.5.0.jar", "Path to an existing AEM quickstart JAR. For example, the latest AEM Cloud SDK quickstart JAR.")
	setupCmd.Flags().StringP("mount", "m", "", "Path to an existing AEM repository that will be mounted to your instances.")
}
