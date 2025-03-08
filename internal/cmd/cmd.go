package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/braam76/bullsh/internal/lua"
)

func ShellPrompt(conf *lua.Config) {
	for envKey, val := range conf.ExportVars {
		os.Setenv(envKey, val)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s", os.Getenv("PS1"))

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			continue
		}

		if err := execInput(input, conf.Aliases); err != nil {
			log.Println(err)
		}
	}
}

func execInput(input string, aliases map[string]string) error {
	input = strings.TrimSpace(input)

	if input == "" {
		return nil
	}

	args := strings.Split(input, " ")

	if command, exists := aliases[args[0]]; exists {
		// Split the alias command into arguments
		aliasArgs := strings.Fields(command)
		args = append([]string{aliasArgs[0]}, aliasArgs[1:]...) // replace args with alias command
	}

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

func cd(args []string) error {
	if len(args) == 1 {
		return os.Chdir(os.Getenv("HOME"))
	}
	return os.Chdir(args[1])
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
