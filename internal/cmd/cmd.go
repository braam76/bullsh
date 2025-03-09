package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func ShellPrompt() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s", os.Getenv("PS1"))

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			continue
		}

		if err := execInput(input); err != nil {
			log.Println(err)
		}
	}
}

func execInput(input string) error {
	input = strings.TrimSpace(input)

	if input == "" {
		return nil
	}

	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		return cd(args)
	case "exit":
		os.Exit(0)
	default:
		return runCommand(args)
	}

	return nil
}

func runCommand(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no command provided")
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
