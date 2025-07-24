package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func cmdClear(args []string, config *Config) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Unable to clear Screen: %v", err)
	}
	return nil
}
