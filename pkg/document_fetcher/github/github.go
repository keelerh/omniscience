package github

import (
	"context"
	"io"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/go-github/github"
	pb "github.com/keelerh/omniscience/protos"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

const (
	reposToListPerPage = 25 // Maximum number of Github repos to return per page.
	teamsToListPerPage = 25 // Maximum number of Github teams to return per page.
	service            = "github"
)

type GithubService struct {
	client         *github.Client
	ingesterClient pb.IngesterClient
	organization   string
}

func NewGithub(token string, organization string, ingesterClient *pb.IngesterClient) *GithubService {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	client := github.NewClient(tc)

	return &GithubService{
		client:         client,
		ingesterClient: *ingesterClient,
		organization:   organization,
	}
}

func (g *GithubService) Fetch(modifiedSince time.Time) error {
	modifiedSinceProto, err := ptypes.TimestampProto(modifiedSince)
	if err != nil {
		return errors.Wrap(err, "failed to parse modified since timestamp as proto")
	}

	ctx := context.Background()
	stream, err := g.ingesterClient.Ingest(ctx)
	if err != nil {
		return err
	}

	repos, err := g.getReposInOrganization(ctx)
	if err != nil {
		return err
	}
	for _, r := range repos {
		owner := r.GetOwner().GetLogin()
		readme, _, err := g.client.Repositories.GetReadme(ctx, owner, r.GetName(), nil)
		if err != nil {
			return err
		}
		content, err := readme.GetContent()
		if err != nil {
			return err
		}
		// Individual files do not have an ID but as we are currently only retrieving
		// the preferred README for a repository we can instead use the repo ID.
		id := strconv.FormatInt(r.GetID(), 10)
		doc := pb.Document{
			Id:           &pb.DocumentId{Id: id},
			Title:        r.GetName() + "/" + readme.GetName(),
			Description:  "",
			Service:      service,
			Content:      content,
			Url:          readme.GetHTMLURL(),
			LastModified: modifiedSinceProto,
		}
		if err := stream.Send(&doc); err != nil {
			return err
		}
	}

	if _, err := stream.CloseAndRecv(); err != nil {
		// We expect io.EOF once the stream has closed.
		if err != io.EOF {
			return err
		}
	}

	return nil
}

func (g *GithubService) getReposInOrganization(ctx context.Context) ([]*github.Repository, error) {
	teams, err := g.getTeamsInOrganization(ctx)
	if err != nil {
		return nil, err
	}

	var repos []*github.Repository
	for _, t := range teams {
		rs, err := g.getReposInTeam(ctx, *t.ID)
		if err != nil {
			return nil, err
		}
		repos = append(repos, rs...)
	}

	return repos, nil
}

func (g *GithubService) getTeamsInOrganization(ctx context.Context) ([]*github.Team, error) {
	var teams []*github.Team
	listOpts := &github.ListOptions{
		PerPage: teamsToListPerPage,
	}
	pageIdx := 0
	for {
		listOpts.Page = pageIdx
		ts, _, err := g.client.Organizations.ListTeams(ctx, g.organization, listOpts)
		if err != nil {
			return nil, err
		}
		teams = append(teams, ts...)
		if len(ts) == teamsToListPerPage {
			pageIdx += teamsToListPerPage
		} else {
			break
		}
	}

	return teams, nil
}

func (g *GithubService) getReposInTeam(ctx context.Context, teamId int64) ([]*github.Repository, error) {
	var repos []*github.Repository
	listOpts := &github.ListOptions{
		PerPage: reposToListPerPage,
	}
	pageIdx := 0
	for {
		listOpts.Page = pageIdx
		rs, _, err := g.client.Organizations.ListTeamRepos(ctx, teamId, listOpts)
		if err != nil {
			return nil, err
		}
		repos = append(repos, rs...)
		if len(rs) == reposToListPerPage {
			pageIdx += reposToListPerPage
		} else {
			break
		}
	}

	return repos, nil
}
