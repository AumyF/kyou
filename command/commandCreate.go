package command

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

var commandCreate = &cli.Command{
	Name:   "create",
	Usage:  "Create a directory inside today's working directory",
	Action: actionCreate,
}

func actionCreate(c *cli.Context) error {
	dirTitle := c.Args().First()

	todaysDir, err := getTodaysDir()
	if err != nil {
		return err
	}

	if dirTitle == "" {
		p := promptui.Prompt{
			Label: "Directory name",
			Validate: func(s string) error {
				if !fs.ValidPath(s) {
					return errors.New("invalid path")
				}
				return nil
			},
		}

		result, err := p.Run()

		if err != nil {
			return fmt.Errorf("prompt failed %v", err)
		}

		dirTitle = result
	}
	dir := filepath.Join(todaysDir, dirTitle)

	fmt.Printf("creating %v\n", dir)
	err = os.MkdirAll(dir, 0o755)
	if err != nil {
		return err
	}

	return nil

}
