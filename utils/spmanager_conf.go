package utils

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type SpmanagerConf struct {
	SpmanagerKey string `yaml:"spmanagerKey"`
	SpmanagerValue string `yaml:"spmanagerValue"`
}

func LoadSpmanagerConf(filePath string) []SpmanagerConf {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Failed to read spmanager config file, " + err.Error())
	}

	var spmanagerConf []SpmanagerConf
	err = yaml.Unmarshal(file, &spmanagerConf)
	if err != nil {
		fmt.Println("Failed to unmarshal spmanager config file, please check: " + err.Error())
	}

	return spmanagerConf
}

func WriteSpmanagerConf(filePath string, spmanagerConf []SpmanagerConf) {
	file, err := yaml.Marshal(spmanagerConf)
	if err != nil {
		fmt.Println("Failed to marshal spmanager config file, " + err.Error())
	}

	err = os.WriteFile(filePath, file, os.ModeAppend)
	if err != nil {
		fmt.Println("Failed to write spmanager config file, " + err.Error())
	}
	fmt.Println("write spmanager config file successfully.")
}