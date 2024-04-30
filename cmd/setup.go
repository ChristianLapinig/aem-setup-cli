/*
Copyright © 2024 Christian Lapinig <lapinig.a.christian@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ChristianLapinig/aem-setup-cli/lib"
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

		if len(args) < 2 {
			fmt.Println("Path to quickstart JAR is required. Exiting...")
			os.Exit(1)
		}

		licenseProperties := args[0]
		quickstartPath := args[1]
		setupPath, _ := cmd.Flags().GetString("path")
		repoPath, _ := cmd.Flags().GetString("mount")
		instance, _ := cmd.Flags().GetString("instance")

		authorPath := filepath.Join(setupPath, "author")
		publishPath := filepath.Join(setupPath, "publish")

		fmt.Println("Creating author and publish folders in", setupPath)
		lib.CheckErr(lib.CreateFolder(authorPath))
		lib.CheckErr(lib.CreateFolder(publishPath))
		fmt.Println("Successfully created author and publish folders")

		fmt.Println("Copying license.properties file to respective instances")
		lib.CheckErr(lib.Copy(licenseProperties, filepath.Join(authorPath, "license.properties")))
		lib.CheckErr(lib.Copy(licenseProperties, filepath.Join(publishPath, "license.properties")))
		fmt.Println("Successfully copied license.properties to author and publish folders")

		// TODO: It should be specified where to mount the given repository
		// should it be in the author or publish folder? or both?

		fmt.Println("Copying quickstart JAR to respective folders")
		lib.CheckErr(lib.Copy(quickstartPath, filepath.Join(authorPath, "aem-author-p4502.jar")))
		lib.CheckErr(lib.Copy(quickstartPath, filepath.Join(publishPath, "aem-publish-p4503.jar")))
		fmt.Println("Successfully copied the quickstart jar to author and publish folder")

		if repoPath != "" && instance != "" {
			fmt.Println("Copying existing repository to", setupPath+"/"+instance)
			lib.CheckErr(lib.Copy(repoPath, filepath.Join(setupPath, instance, "crx-quickstart")))
			fmt.Println("Successfully mounted existing repository to", setupPath+"/"+instance)
		} else if repoPath != "" && instance == "" {
			fmt.Println("An instance must be specified when using the 'mount' command.")
			os.Exit(1)
		}

		fmt.Println("Done. Local AEM environment is now setup.")
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
	setupCmd.Flags().StringP("path", "p", ".", "Path where AEM should be setup. Default is the current directory you are in.")
	setupCmd.Flags().StringP("mount", "m", "", "Path to an existing AEM repository that will be mounted to your instances. Must be used with the instance flag.")
	setupCmd.Flags().StringP("instance", "i", "", "Instance where the given repository from 'mount' should be mounted.")
}
