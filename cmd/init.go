/*
Copyright © 2024 shaoyang

*/
package cmd

import (
	"fmt"
	"github.com/FDU-YSP/spmanager/utils"
	"github.com/spf13/cobra"
	"path/filepath"
)

const (
	spmanagerFolder  = ".spmanager"
	spmanagerCfgYaml = "spmanager.conf"
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
		if homeDir != "" {
			// create .spmanager folder
			utils.CreateDir(filepath.Join(homeDir, spmanagerFolder))

			// create spmanager.conf file
			if !utils.CheckFile(filepath.Join(homeDir, spmanagerFolder, spmanagerCfgYaml)) {
				// create spmanager.conf file
				utils.CreateFile(filepath.Join(homeDir, spmanagerFolder, spmanagerCfgYaml))
				fmt.Println("spmanager init succeed.")
			} else {
				fmt.Println("spmanager.conf already exists, will skip.")
			}

		} else {
			fmt.Println("spmanager cannot find home directory, please check your user home path on system")
			fmt.Println("spmanager init failed.")
		}

	},
}

func init() {

	rootCmd.AddCommand(initCmd)
	
}
