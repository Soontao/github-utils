package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"
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

		csvOutputFile, err := os.OpenFile(output, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			return err
		}

		csvWriter := csv.NewWriter(csvOutputFile)
		defer csvOutputFile.Close()
		defer csvWriter.Flush()

		csvEncoder := csvutil.NewEncoder(csvWriter)

		if err := csvEncoder.EncodeHeader(ExportIssue{}); err != nil {
			return err
		}

		total := 0

		log.Println("Exporting...")
		defer log.Println("Finished")

		return lib.RunWithAllIssuesInOrg(client, ctx, org, func(repo *github.Repository, issue *github.Issue) error {

			labels := []string{}
			for _, label := range issue.Labels {
				labels = append(labels, label.GetName())
			}

			// not pull request
			if issue.GetPullRequestLinks() == nil {
				total++
				if err := csvEncoder.Encode(ExportIssue{
					repo.GetFullName(),
					issue.GetTitle(),
					*issue.User.Login,
					issue.CreatedAt,
					issue.GetState(),
					strings.Join(labels, ","),
					issue.GetHTMLURL(),
				}); err != nil {
					return err
				}

			}

			return nil
		})

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
