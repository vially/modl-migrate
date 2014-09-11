package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	ui := &cli.BasicUi{Writer: os.Stdout}

	cli := &cli.CLI{
		Args: os.Args[1:],
		Commands: map[string]cli.CommandFactory{
			"up": func() (cli.Command, error) {
				return &UpCommand{
					Ui: ui,
				}, nil
			},
		},
		HelpFunc: cli.BasicHelpFunc("gorp-migrate"),
		Version:  "1.0.0",
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitCode
}
