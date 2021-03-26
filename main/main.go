package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli"
)

// Version string, in release version
// This variable will be overwrited by complier
var Version = "SNAPSHOT"

// AppName of this application
var AppName = "GithubUtil"

// AppUsage of this application
var AppUsage = "Github Util Tool"

func main() {
	app := cli.NewApp()
	app.Version = Version
	app.Name = AppName
	app.Usage = AppUsage
	app.Flags = options
	app.EnableBashCompletion = true

	commonCommands := []cli.Command{
		commandExport,
	}

	app.Commands = commonCommands

	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
