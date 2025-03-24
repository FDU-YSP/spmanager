/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/FDU-YSP/spmanager/utils"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of spmanager",
	Long: `Print the version number of spmanager:

You can use 'spmanager version'.`,
	Run: func(cmd *cobra.Command, args []string) {
		ver := utils.GetVersion()
		fmt.Printf("spmanager version %s\n", ver)
	},
}

func init() {

	rootCmd.AddCommand(versionCmd)
}
