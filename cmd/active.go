/*
Copyright © 2024 shaoyang
*/
package cmd

import (
	"fmt"
	"github.com/FDU-YSP/spmanager/utils"

	"github.com/spf13/cobra"
)

// activeCmd represents the active command
var activeCmd = &cobra.Command{
	Use:   "active",
	Short: "active the shell property to your current shell",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("active called")

		allFlags := cmd.Flags()
		key, _ := allFlags.GetString("key")

		cfg := utils.LoadSpmanagerConf(utils.GetSpmanagerConfigFile())

		var resultSet []utils.SpmanagerProperties
		for _, prop := range cfg.Properties {
			if prop.SpmanagerKey == key {
				resultSet = append(resultSet, prop)
			}
		}
		if len(resultSet) == 0 {
			fmt.Println("No such shell property found in your spmanager.")
		} else if len(resultSet) == 1 {
			fmt.Println("The shell property you want to execute is: ", resultSet[0].SpmanagerValue)

			// Execute the shell command
			cmd := resultSet[0].SpmanagerValue
			utils.SourceShellCommand(cmd)
		} else {
			fmt.Println("Multiple shell properties found in your spmanager, please correct those duplicates.")
		}
	},
}

func init() {

	rootCmd.Flags().StringP("key", "k", "", "the shell command you want to execute.")

	_ = addCmd.MarkFlagRequired("key")

	rootCmd.AddCommand(activeCmd)
}
