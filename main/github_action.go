package main

import (
	"context"

	"github.com/google/go-github/v34/github"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
)

type GithubAction func(client *github.Client, ctx context.Context, cli *cli.Context) error

func toCliAction(ghAction GithubAction) cli.ActionFunc {
	ctx := context.Background()
	return func(c *cli.Context) error {
		gh := c.GlobalString("gh")
		token := c.GlobalString("token")
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		var client *github.Client
		if len(gh) > 0 {
			c, err := github.NewEnterpriseClient(gh, gh, tc)
			if err != nil {
				return err
			}
			client = c
		} else {
			client = github.NewClient(tc)
		}
		return ghAction(client, ctx, c)
	}
}
