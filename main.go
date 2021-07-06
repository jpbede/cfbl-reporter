package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"go.bnck.me/cfbl-reporter/internal/commands"
	"log"
	"os"

	_ "go.bnck.me/cfbl-reporter/internal/commands/send"
)

var (
	version = "dev"
	commit  = "none"
)

func main() {
	app := &cli.App{
		Name:     "cfbl-reporter",
		Usage:    "",
		Version:  fmt.Sprintf("%s-%s", version, commit),
		Commands: commands.Get(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				DefaultText: "/etc/cfbl-reporter.yaml",
				Value:       "/etc/cfbl-reporter.yaml",
			},
		},
	}

	// run app
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
