/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
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

		if err := cmd.ParseFlags(args); err != nil {
			f := cmd.PersistentFlags()
			key := f.Lookup("key")
			val := f.Lookup("val")
			fmt.Println(homedir.Dir())

			if key == nil || val == nil {
				fmt.Println("key or val cannot be empty")
				return
			}

		} else {
			fmt.Println(err.Error())
			fmt.Println("parse flag error")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.PersistentFlags().StringP("key", "k", "", "linux shell you want to execute.")

	addCmd.PersistentFlags().StringP("val", "v", "", "your customized shell content.")

	addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
