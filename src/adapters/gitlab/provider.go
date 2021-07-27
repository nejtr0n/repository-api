package gitlab

import (
	"crypto/tls"
	"net/http"

	"github.com/nejtr0n/repository-api/src/domain"

	"github.com/xanzy/go-gitlab"
)

func NewProvider(token domain.GitlabApiToken, baseUrl domain.GitlabBaseUrl, vcs domain.GitlabVcs) (*Provider, error) {
	client, err := gitlab.NewClient(token.String(),
		gitlab.WithBaseURL(baseUrl.String()),
		gitlab.WithHTTPClient(buildHttpClient()),
	)
	if err != nil {
		return nil, err
	}
	return &Provider{
		client: client,
		vcs:    vcs,
	}, nil
}

type Provider struct {
	client *gitlab.Client
	vcs    domain.Vcs
}

func (p Provider) GetProjectById(id int) (*domain.Project, error) {
	project, _, err := p.client.Projects.GetProject(id, nil)
	if err != nil {
		return nil, err
	}
	return &domain.Project{
		Id:       project.ID,
		Name:     project.Name,
		Provider: p,
		VcsUrl:   project.HTTPURLToRepo,
	}, nil
}

func (p Provider) GetType() domain.ProviderType {
	return domain.ProviderGitlab
}

func (p Provider) GetVcs() domain.Vcs {
	return p.vcs
}

func (p Provider) CreateProject(name string, groupId int) (*domain.Project, error) {
	opts := &gitlab.CreateProjectOptions{
		Name:                 gitlab.String(name),
		MergeRequestsEnabled: gitlab.Bool(true),
		SnippetsEnabled:      gitlab.Bool(false),
		Visibility:           gitlab.Visibility(gitlab.PublicVisibility),
	}
	if groupId > 0 {
		opts.NamespaceID = &groupId
	}

	project, _, err := p.client.Projects.CreateProject(opts)
	if err != nil {
		return nil, err
	}

	return &domain.Project{
		Id:       project.ID,
		Name:     project.Name,
		Provider: p,
		VcsUrl:   project.HTTPURLToRepo,
	}, nil
}

func buildHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}
