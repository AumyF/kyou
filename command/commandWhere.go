package command

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var commandWhere = &cli.Command{
	Name:   "where",
	Usage:  "Print absolute path of today's working directory",
	Action: actionWhere,
}

func actionWhere(c *cli.Context) error {
	todaysDir, err := getTodaysDir()
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", todaysDir)

	return nil
}
