/*
Copyright © 2024 shaoyang
*/
package cmd

import (
	"fmt"
	"github.com/FDU-YSP/spmanager/utils"
	"github.com/spf13/cobra"
)

const spmanagerCfg = ""

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add shell property to spmanager",
	Long: `Add your shell properties to spmanager.

so far, spmanager only support "alias", "source" and "export" CLI`,
	Run: func(cmd *cobra.Command, args []string) {

		allFlags := cmd.Flags()
		key, _ := allFlags.GetString("key")
		val, _ := allFlags.GetString("val")
		if spmCfg := utils.GetSpmanagerConfigFile(); spmCfg != "" {
			fmt.Println("add shell property to spmanager")

			// read spmanager config file(yaml file)
			spmanagerCfg := utils.LoadSpmanagerConf(spmCfg)
			props := spmanagerCfg.Properties
			props = append(props, utils.SpmanagerProperties{
				SpmanagerKey:   key,
				SpmanagerValue: val,
			})

			spmanagerCfg.Properties = props
			spmanagerCfg.Size = len(props)

			// write spmanager config file
			utils.WriteSpmanagerConf(spmCfg, spmanagerCfg)
		} else {
			fmt.Println("cannot find spmanager config file, please check your config file.")
		}

	},
}

func init() {

	addCmd.Flags().StringP("key", "k", "", "the shell command you want to execute.")
	addCmd.Flags().StringP("val", "v", "", "your customized shell content.")
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	_ = addCmd.MarkFlagRequired("key")
	_ = addCmd.MarkFlagRequired("val")

	// must put this line after addCmd.Flags()
	rootCmd.AddCommand(addCmd)
}
