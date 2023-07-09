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
	Owner          string
	Name           string
	FullName       string
	Description    string
	RepoURL        string
	DefaultBranch  string
	Parent         string
	ParentFullName string
	ParentDeleted  bool
	Private        bool
	BehindBy       int
}

func GetForks(ctx context.Context, client *github.Client) ([]*RepositoryWithDetails, error) {
	var forksWithDetails []*RepositoryWithDetails
	forks, err := getAllForks(context.Background(), client)

	if err != nil {
		return forksWithDetails, fmt.Errorf("failed to fetch fork list: %w\n", err)
	}

	for _, fork := range forks {

		repo, resp, err := client.Repositories.Get(ctx, fork.GetOwner().GetLogin(), fork.GetName())

		switch resp.StatusCode {
		case http.StatusForbidden:
			continue
		case http.StatusUnavailableForLegalReasons:
			forksWithDetails = append(forksWithDetails, buildDetails(fork, nil, resp.StatusCode))
			continue
		}

		if err != nil {
			return forksWithDetails, fmt.Errorf("failed to get repository %s: %w", fork.GetName(), err)
		}

		parent := repo.GetParent()

		base := fmt.Sprintf("%s:%s", parent.GetOwner().GetLogin(), repo.GetDefaultBranch()) // compare with forked repo's default branch
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
			return forksWithDetails, fmt.Errorf("failed to compare repository with parent %s: %w", parent.GetName(), err)
		}

		if cmpr.GetBehindBy() < 1 {
			continue
		}

		forksWithDetails = append(forksWithDetails, buildDetails(repo, cmpr, resp.StatusCode))

	}

	sort.SliceStable(forksWithDetails, func(i, j int) bool {
		return forksWithDetails[i].BehindBy > forksWithDetails[j].BehindBy
	})

	return forksWithDetails, nil
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
	if repo == nil || commit == nil {
		return &RepositoryWithDetails{}
	}

	return &RepositoryWithDetails{
		Owner:          repo.GetOwner().GetLogin(),
		Name:           repo.GetName(),
		FullName:       repo.GetFullName(),
		Description:    repo.GetDescription(),
		RepoURL:        repo.GetURL(),
		DefaultBranch:  repo.GetDefaultBranch(),
		Parent:         repo.GetParent().GetOwner().GetLogin(),
		ParentFullName: repo.GetParent().GetFullName(),
		ParentDeleted:  code == http.StatusNotFound,
		Private:        repo.GetPrivate(),
		BehindBy:       commit.GetBehindBy(),
	}
}

func getAllForks(ctx context.Context, client *github.Client) ([]*github.Repository, error) {
	var allRepos []*github.Repository
	opts := &github.RepositoryListOptions{
		Type:        "owner",
		ListOptions: github.ListOptions{PerPage: pageSize},
	}

	for {
		forks, resp, err := client.Repositories.List(ctx, "", opts)
		if err != nil {
			return allRepos, err
		}

		allRepos = append(allRepos, forks...)

		if resp.NextPage == 0 {
			break
		}

		opts.Page = resp.NextPage
	}

	forks := make([]*github.Repository, 0, len(allRepos))
	for _, repo := range allRepos {
		if !repo.GetFork() {
			continue
		}

		forks = append(forks, repo)
	}

	return forks, nil
}
