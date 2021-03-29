package main

import "github.com/urfave/cli"

var options = []cli.Flag{
	cli.StringFlag{
		Name:     "token, t",
		EnvVar:   "GH_TOKEN",
		Usage:    "Github Personal Access Token",
		Required: true,
	},
	cli.StringFlag{
		Name:     "gh, g",
		Usage:    "Github Instance",
		EnvVar:   "GH_INSTANCE",
		Required: false,
	},
}
