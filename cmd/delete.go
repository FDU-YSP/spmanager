/*
Copyright © 2024 shaoyang

*/
package cmd

import (
	"fmt"
	"github.com/FDU-YSP/spmanager/utils"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete shell property from spmanager",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		allFlags := cmd.Flags()
		key, _ := allFlags.GetBool("all")
		if key {
			// delete all properties
			spmanagerCfg := utils.LoadSpmanagerConf(utils.GetSpmanagerConfigFile())
			spmanagerCfg.Size = 0
			spmanagerCfg.Properties = make([]utils.SpmanagerProperties, 0)

			// write back to the config file
			utils.WriteSpmanagerConf(utils.GetSpmanagerConfigFile(), spmanagerCfg)
		} else {
			// delete a specific property
			key, _ := allFlags.GetString("key")
			spmanagerCfg := utils.LoadSpmanagerConf(utils.GetSpmanagerConfigFile())

			var resultSet []utils.SpmanagerProperties
			for _, prop := range spmanagerCfg.Properties {
				if prop.SpmanagerKey == key {
					resultSet = append(resultSet, prop)
				}
			}
			if len(resultSet) == 0 {
				fmt.Println("No such shell property found in your spmanager.")
			} else if len(resultSet) == 1 {
				fmt.Println("The shell property you want to delete is: ", resultSet[0].SpmanagerKey)
				// delete the shell property
				var newProperties []utils.SpmanagerProperties
				for _, prop := range spmanagerCfg.Properties {
					if prop.SpmanagerKey != key {
						newProperties = append(newProperties, prop)
					}
				}
				spmanagerCfg.Properties = newProperties
				spmanagerCfg.Size = len(newProperties)

				// write back to the config file
				utils.WriteSpmanagerConf(utils.GetSpmanagerConfigFile(), spmanagerCfg)
			} else {
				fmt.Println("Multiple same shell properties found in your spmanager, please correct those duplicates.")
			}
		}
	},
}

func init() {

	deleteCmd.Flags().StringP("key", "k", "", "key of the property to delete")
	deleteCmd.Flags().BoolP("all", "a", false, "delete all properties")

	// _ = deleteCmd.MarkFlagRequired("key")

	rootCmd.AddCommand(deleteCmd)
}