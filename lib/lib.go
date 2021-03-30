package lib

import (
	"context"
	"log"

	"github.com/fatih/color"
	"github.com/google/go-github/v34/github"
	"golang.org/x/sync/semaphore"
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
	maxConcurrency := int64(10)
	rContext := context.Background()
	sem := semaphore.NewWeighted(maxConcurrency)

	err := RunWithAllRepoInOrgs(client, ctx, org, func(repo *github.Repository) error {

		sem.Acquire(rContext, 1)

		go func() {

			defer sem.Release(1)
			pageIdx := 1
			repoName := repo.GetName()

			for {
				issues, resp, err := client.Issues.ListByRepo(ctx, org, repoName, &github.IssueListByRepoOptions{
					ListOptions: github.ListOptions{Page: pageIdx, PerPage: 100},
					State:       "all",
				})
				if err != nil {
					log.Println(color.RedString("error: ", err))
					return
				}
				for _, issue := range issues {
					if err := consumer(repo, issue); err != nil {
						log.Println(color.RedString("error: ", err))
						return
					}
				}
				if resp.NextPage == 0 {
					break
				}
				pageIdx++
			}

		}()

		return nil

	})

	if err != nil {
		return err
	}

	// wait finished
	sem.Acquire(rContext, maxConcurrency)

	return nil

}
