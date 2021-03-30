package lib

import (
	"context"
	"log"

	"github.com/fatih/color"
	"github.com/google/go-github/v34/github"
)

type OrgConsumer func(*github.Repository) error
type IssueConsumer func(*github.Repository, *github.Issue) error

func RunWithAllRepoInOrgs(client *github.Client, ctx context.Context, org string, consumer OrgConsumer) error {
	pageIdx := 1
	for {
		repos, resp, err := client.Repositories.ListByOrg(
			ctx,
			org,
			&github.RepositoryListByOrgOptions{ListOptions: github.ListOptions{Page: pageIdx, PerPage: 100}},
		)
		if err != nil {
			return err
		}
		for _, repo := range repos {
			if err := consumer(repo); err != nil {
				return err
			}
		}
		if resp.NextPage == 0 {
			break
		}
		pageIdx++
	}
	return nil
}

func RunWithAllIssuesInOrg(client *github.Client, ctx context.Context, org string, consumer IssueConsumer) error {

	return RunWithAllRepoInOrgs(client, ctx, org, func(repo *github.Repository) error {

		pageIdx := 1

		repoName := repo.GetName()

		for {
			issues, resp, err := client.Issues.ListByRepo(ctx, org, repoName, &github.IssueListByRepoOptions{
				ListOptions: github.ListOptions{Page: pageIdx, PerPage: 100},
				State:       "all",
			})
			if err != nil {
				return err
			}
			log.Println(color.GreenString("Got %v/%v/issues?page=%v", org, repoName, pageIdx))
			for _, issue := range issues {
				if err := consumer(repo, issue); err != nil {
					return err
				}
			}
			if resp.NextPage == 0 {
				break
			}
			pageIdx++
		}

		return nil
	})

}
