package domain

import (
	"errors"

	"github.com/urfave/cli/v2"
)

const (
	ProviderGitlab = ProviderType("gitlab")
)

type GitlabApiToken string

func (g GitlabApiToken) String() string {
	return string(g)
}

func NewGitlabApiToken(config *cli.Context) GitlabApiToken {
	return GitlabApiToken(config.String("gitlab_token"))
}

type GitlabBaseUrl string

func (g GitlabBaseUrl) String() string {
	return string(g)
}

func NewGitlabBaseUrl(config *cli.Context) GitlabBaseUrl {
	return GitlabBaseUrl(config.String("gitlab_base_url"))
}

var (
	UnsupportedProviderError = errors.New("unsupported provider")
)

type ProviderType string

func (p ProviderType) String() string {
	return string(p)
}

type Provider interface {
	GetType() ProviderType
	GetVcs() Vcs
	CreateProject(name string, groupId int) (*Project, error)
	GetProjectById(id int) (*Project, error)
}

type GitlabProvider interface {
	Provider
}

type ProviderFactory interface {
	GetProvider(providerType ProviderType) (Provider, error)
}

func NewProviderFactory(gitlab GitlabProvider) ProviderFactory {
	return providerFactory{
		gitlab: gitlab,
	}
}

type providerFactory struct {
	gitlab Provider
}

func (p providerFactory) GetProvider(providerType ProviderType) (Provider, error) {
	switch providerType {
	case ProviderGitlab:
		return p.gitlab, nil
	default:
		return nil, UnsupportedProviderError
	}
}
