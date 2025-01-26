/*
Copyright © 2024 shaoyang

*/
package cmd

import (
	"fmt"
	"github.com/FDU-YSP/spmanager/utils"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list all the shell profiles in your spmanager:")
		cfg := utils.LoadSpmanagerConf(utils.GetSpmanagerConfigFile())

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Property Name", "Property Value"})

		for i, prop := range cfg {
			table.Append([]string{strconv.Itoa(i+1), prop.SpmanagerKey, prop.SpmanagerValue})
		}

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
