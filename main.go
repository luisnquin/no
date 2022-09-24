package main

import (
	"bufio"
	_ "embed"
	"os"

	"github.com/urfave/cli/v2"
)

//go:embed help.tpl
var helpTemplate string

func main() {
	app := cli.App{
		Name:        "no",
		Description: "Repeatedly output a line with all specified STRING(s), or 'n'. Antithesis of 'yes'.",
		Authors:     []*cli.Author{{Name: "luisnquin", Email: "lpaadres2020@gmail.com"}},
		Version:     "v0.0.1",
		Flags: []cli.Flag{
			cli.VersionFlag,
			cli.HelpFlag,
		},
		CustomAppHelpTemplate: helpTemplate,
		Action: func(ctx *cli.Context) error {
			msgs := []string{"n"}

			if len(os.Args) > 1 {
				msgs = os.Args[1:]
			}

			w := bufio.NewWriter(os.Stdout)

			for {
				for _, msg := range msgs {
					w.WriteString(msg)
					w.WriteString(" ")
				}

				w.WriteString("\n")
			}
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
