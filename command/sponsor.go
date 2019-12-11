package command

import (
	"github.com/urfave/cli/v2"
	"insecurecert/utils"
	"log"
)

type Sponsor struct{}

func (Sponsor) Command() *cli.Command {
	return &cli.Command{
		Name:    "sponsor",
		Aliases: []string{"donate", "donation"},
		Usage:   "opens the sponsorship URL in browser",
		Action: func(c *cli.Context) error {
			log.Print("Do you know GitHub is going to match your sponsorship for a year?")
			log.Print("10x developer is a myth but 2x sponsor is real")

			return utils.Browser{}.Open("https://github.com/sponsors/TangRufus")
		},
	}
}
