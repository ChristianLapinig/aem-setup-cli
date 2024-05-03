/*
Copyright © 2024 Christian Lapinig <lapinig.a.christian@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aem-setup-cli",
	Short: "CLI tool to help setup AEM locally",
	Long:  `CLI tool that makes it easy to setup your environment to run AEM locally.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	setup := NewSetupImpl()
	setupCmd := NewSetupCommand(setup)

	rootCmd.AddCommand(setupCmd)
}
