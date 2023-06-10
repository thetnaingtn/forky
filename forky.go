package forky

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/google/go-github/v52/github"
)

var (
	pageSize = 100
)

type RepositoryWithDetails struct {
	Owner         string
	Name          string
	FullName      string
	Description   string
	RepoURL       string
	DefaultBranch string
	Parent        string
	ParentDeleted bool
	Private       bool
	BehindBy      int
}

func GetForks(ctx context.Context, client *github.Client) ([]*RepositoryWithDetails, error) {
	var forks []*RepositoryWithDetails
	repos, err := getAllRepositories(context.Background(), client)

	if err != nil {
		return forks, err
	}

	for _, r := range repos {
		if !r.GetFork() {
			continue
		}
		repo, resp, err := client.Repositories.Get(ctx, r.GetOwner().GetLogin(), r.GetName())

		switch resp.StatusCode {
		case http.StatusForbidden:
			continue
		case http.StatusUnavailableForLegalReasons:
			forks = append(forks, buildDetails(r, nil, resp.StatusCode))
			continue
		}

		if err != nil {
			return forks, fmt.Errorf("failed to get repository %s: %w", r.GetName(), err)
		}

		parent := repo.GetParent()

		base := fmt.Sprintf("%s:%s", parent.GetOwner().GetLogin(), parent.GetDefaultBranch())
		head := fmt.Sprintf("%s:%s", repo.GetOwner().GetLogin(), repo.GetDefaultBranch())
		cmpr, resp, err := client.Repositories.CompareCommits(
			ctx,
			repo.GetOwner().GetLogin(),
			repo.GetName(),
			base,
			head,
			&github.ListOptions{},
		)

		if err != nil && resp.StatusCode != http.StatusNotFound {
			log.Println("ERR", err)
			return forks, fmt.Errorf("failed to compare repository with parent %s: %w", parent.GetName(), err)
		}

		if cmpr == nil || cmpr.GetBehindBy() < 1 {
			continue
		}

		forks = append(forks, buildDetails(repo, cmpr, resp.StatusCode))

	}

	sort.SliceStable(forks, func(i, j int) bool {
		return forks[i].BehindBy > forks[j].BehindBy
	})

	return forks, nil
}

func SyncBranchWithUpstreamRepo(client *github.Client, repo *RepositoryWithDetails) error {
	request := &github.RepoMergeUpstreamRequest{Branch: &repo.DefaultBranch}
	res, resp, err := client.Repositories.MergeUpstream(context.Background(), repo.Owner, repo.Name, request)

	if resp.StatusCode == http.StatusConflict {
		return fmt.Errorf("couldn't merge with upstream %s branch due to conflict", res.GetBaseBranch())
	}

	if err != nil {
		return fmt.Errorf("couldn't merge with upstream %s branch: %w", res.GetBaseBranch(), err)
	}

	return nil
}

func buildDetails(repo *github.Repository, commit *github.CommitsComparison, code int) *RepositoryWithDetails {
	return &RepositoryWithDetails{
		Owner:         repo.GetOwner().GetLogin(),
		Name:          repo.GetName(),
		FullName:      repo.GetFullName(),
		Description:   repo.GetDescription(),
		RepoURL:       repo.GetURL(),
		DefaultBranch: repo.GetDefaultBranch(),
		Parent:        repo.GetParent().GetFullName(),
		ParentDeleted: code == http.StatusNotFound,
		Private:       repo.GetPrivate(),
		BehindBy:      commit.GetBehindBy(),
	}
}

func getAllRepositories(ctx context.Context, client *github.Client) ([]*github.Repository, error) {
	var allRepos []*github.Repository
	opts := &github.RepositoryListOptions{
		Type:        "owner",
		ListOptions: github.ListOptions{PerPage: pageSize},
	}

	for {
		repos, resp, err := client.Repositories.List(ctx, "", opts)
		if err != nil {
			return allRepos, err
		}

		allRepos = append(allRepos, repos...)

		if resp.NextPage == 0 {
			break
		}

		opts.Page = resp.NextPage
	}

	return allRepos, nil
}
