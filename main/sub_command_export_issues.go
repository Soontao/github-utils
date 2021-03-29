package main

import (
	"context"
	"io/ioutil"
	"strings"
	"time"

	"github.com/Soontao/github-utils/lib"
	"github.com/google/go-github/v34/github"
	"github.com/jszwec/csvutil"
	"github.com/urfave/cli"
)

type ExportIssue struct {
	Repository string
	Title      string
	CreatedBy  string
	DateTime   *time.Time
	Status     string
	Labels     string
	Link       string
}

var subCommandExportIssues = cli.Command{
	Name:  "issues",
	Usage: "program entry",
	Action: toCliAction(func(client *github.Client, ctx context.Context, cli *cli.Context) error {

		org := cli.String("org")
		output := cli.String("file")
		exportIssues := []ExportIssue{}

		lib.RunWithAllIssuesInOrg(client, ctx, org, func(repo *github.Repository, issue *github.Issue) error {
			labels := []string{}
			for _, label := range issue.Labels {
				labels = append(labels, label.GetName())
			}

			if issue.GetPullRequestLinks() == nil {
				exportIssues = append(exportIssues, ExportIssue{
					repo.GetFullName(),
					issue.GetTitle(),
					*issue.User.Login,
					issue.CreatedAt,
					issue.GetState(),
					strings.Join(labels, ","),
					issue.GetHTMLURL(),
				})
			}

			return nil
		})
		content, err := csvutil.Marshal(exportIssues)
		if err != nil {
			return nil
		}
		if err := ioutil.WriteFile(output, content, 0644); err != nil {
			return err
		}
		return nil
	}),
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "org, o",
			Usage:    "the target orginization",
			Required: true,
		},
		cli.StringFlag{
			Name:     "file, f",
			Usage:    "output file",
			Required: true,
		},
	},
}
