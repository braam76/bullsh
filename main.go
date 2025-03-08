package main

import (
	"fmt"

	"github.com/braam76/bullsh/internal/cmd"
	"github.com/braam76/bullsh/internal/lua"
)

func main() {
	config := &lua.Config{}
	err := config.LoadConfig("./example/config.lua")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", config)
	cmd.ShellPrompt(config)

}
