/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/FDU-YSP/spmanager/utils"
	"github.com/spf13/cobra"
	"path/filepath"
)

const (
	spmanagerFolder = ".spmanager"
	spmanagerConf = "spmanager.conf"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init spmanager CLI",
	Long: `init CLI for spmanager.

It would create ".spmanager" folder under user home directory, And generate a config file named "spmanager.conf" in it.
"spmanager.conf" use yaml file format.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("spmanager init ...")
		homeDir := utils.HomeDir()
		if homeDir != ""{
			// create .spmanager folder
			utils.CreateDir(filepath.Join(homeDir, spmanagerFolder))

			// create spmanager.conf file
			if utils.CheckFile(filepath.Join(homeDir, spmanagerFolder, spmanagerConf)) {
				// create spmanager.conf file
				utils.CreateFile(filepath.Join(homeDir, spmanagerFolder, spmanagerConf))
				fmt.Println("spmanager init succeed.")
			} else {
				fmt.Println("spmanager.conf already exists, will skip.")
			}

		} else {
			fmt.Println("spmanager cannot find home directory, please check your system")
			fmt.Println("spmanager init failed.")
		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
