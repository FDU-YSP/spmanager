package utils

import (
	"fmt"
	"os/exec"
)

func SourceShellCommand(cmd string) {
	c := exec.Command("source", cmd)
	if out, err := c.Output(); err != nil {
		fmt.Println("Failed to source shell command: ", err)
	} else {
		str := string(out)
		if len(str) == 0 {
			fmt.Println("Source shell command successfully.")
		} else {
			fmt.Println(string(out))
		}
	}
}