package main

import (
	"os"

	"github.com/aumyf/kyou/command"
	"github.com/urfave/cli/v2"
)

func main() {
	err := newApp().Run(os.Args)

	if err != nil {
		os.Exit(1)
	}
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "kyou"
	app.Usage = "Create your today's working directory"
	app.Commands = command.Commands
	app.Action = command.DefaultAction

	return app
}
