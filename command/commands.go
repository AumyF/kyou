package command

import (
	"github.com/urfave/cli/v2"
)

var Commands = []*cli.Command{
	commandCreate, commandWhere,
}

var DefaultAction = actionCreate
