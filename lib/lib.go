package lib

import (
	"context"

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

		for {
			pageIdx := 1
			issues, resp, err := client.Issues.ListByRepo(ctx, org, repo.GetName(), &github.IssueListByRepoOptions{
				ListOptions: github.ListOptions{Page: pageIdx, PerPage: 100},
				State:       "all",
			})
			if err != nil {
				return err
			}
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
