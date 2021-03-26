package main

import "github.com/urfave/cli"

var options = []cli.Flag{
	cli.StringFlag{
		Name:     "username, u",
		Usage:    "Github Username",
		Required: true,
	},
	cli.StringFlag{
		Name:     "token, t",
		Usage:    "Github Personal Access Token",
		Required: true,
	},
	cli.StringFlag{
		Name:  "api, a",
		Usage: "Github API Endpoint",
		Value: "https://api.github.com",
	},
}
