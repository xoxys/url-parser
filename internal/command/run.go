package command

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Run default command and print out full url
func Run(ctx *cli.Context) error {
	parts := parseURL(ctx.String("url"))

	if len(parts.String()) > 0 {
		fmt.Println(parts)
	}
	return nil
}
