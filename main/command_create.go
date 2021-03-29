package main

import "github.com/urfave/cli"

var commandCreate = cli.Command{
	Name:        "create",
	Usage:       "Create Something",
	Subcommands: cli.Commands{subCommandCreateLabels},
}
