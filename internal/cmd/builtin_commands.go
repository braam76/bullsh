package cmd

import (
	"fmt"
	"os"
)

func cd(args []string) error {
	if len(args) > 2 {
		fmt.Fprintf(os.Stdout, "USAGE: cd <directory>")
		return nil
	}

	var newDir string

	if len(args) == 1 {
		newDir = os.Getenv("HOME")
	} else {
		newDir = args[1]
	}

	if err := os.Chdir(newDir); err != nil {
		return err
	}

	updatePWD()
	return nil
}

func updatePWD() {
	pwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	os.Setenv("PWD", pwd)
}
