package command

import (
	"github.com/urfave/cli/v2"
	"insecurecert/utils"
)

type Bug struct{}

func (Bug) Command() *cli.Command {
	return &cli.Command{
		Name:    "bug",
		Aliases: []string{"bugs", "issue", "issues"},
		Usage:   "opens bug tracker URL in browser",
		Action: func(c *cli.Context) error {
			return utils.Browser{}.Open("https://github.com/TypistTech/insecurecert/issues")
		},
	}
}
