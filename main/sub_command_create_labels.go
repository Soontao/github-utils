package main

import (
	"context"
	"log"

	"github.com/AvraamMavridis/randomcolor"
	"github.com/fatih/color"
	"github.com/google/go-github/v34/github"
	"github.com/urfave/cli"
)

var subCommandCreateLabels = cli.Command{
	Name:  "labels",
	Usage: "create labels",
	Action: toCliAction(func(client *github.Client, ctx context.Context, cli *cli.Context) error {
		labels := cli.StringSlice("labels")
		org := cli.String("org")
		colors := map[string]string{}
		for _, label := range labels {
			colors[label] = randomcolor.GetRandomColorInHex()[1:]
		}

		pageIdx := 1
		for {
			repos, resp, err := client.Repositories.ListByOrg(
				ctx,
				org,
				&github.RepositoryListByOrgOptions{ListOptions: github.ListOptions{Page: pageIdx}},
			)

			for _, repo := range repos {
				for _, label := range labels {

					c := colors[label]
					_, resp, _ := client.Issues.GetLabel(ctx, org, repo.GetName(), label)
					if resp.StatusCode == 404 {
						log.Println(color.GreenString("creating label '%v' for repo '%v'", label, repo.GetFullName()))
						_, _, err := client.Issues.CreateLabel(ctx, org, repo.GetName(), &github.Label{
							Name:  &label,
							Color: &c,
						})
						if err != nil {
							log.Println(color.RedString("create failed, %v", err))
						}
					} else {
						log.Println(color.YellowString("label '%v' existed in repo '%v'", label, repo.GetFullName()))
					}

				}

			}

			if err != nil {
				return err
			}
			if resp.NextPage == 0 {
				break
			}
			pageIdx++
		}

		return nil
	}),
	Flags: []cli.Flag{
		cli.StringSliceFlag{
			Name:     "labels, l",
			Usage:    "the tags to be created",
			Required: true,
		},
		cli.StringFlag{
			Name:     "org, o",
			Usage:    "the orginazation",
			Required: true,
		},
	},
}
