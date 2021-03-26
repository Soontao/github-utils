package main

import "github.com/urfave/cli"

var subCommandExportIssues = cli.Command{
	Name:  "issues",
	Usage: "program entry",
	Action: func(*cli.Context) error {
		return nil
	},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "username, u",
			Required: true,
		},
		cli.StringFlag{
			Name:     "token, t",
			Required: true,
		},
	},
}