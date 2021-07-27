package git

import (
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5/config"
	githttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/nejtr0n/repository-api/src/domain"

	"github.com/go-git/go-git/v5"
)

var defaultRemoteName = "clone"

func NewVcs(token domain.GitlabApiToken) *Vcs {
	return &Vcs{token: token}
}

type Vcs struct {
	token domain.GitlabApiToken
}

func (v Vcs) CloneRepository(project *domain.Project) (*git.Repository, error) {
	repository, err := git.Clone(memory.NewStorage(), memfs.New(), &git.CloneOptions{
		URL: project.VcsUrl,
		Auth: &githttp.BasicAuth{
			Username: "nobody",
			Password: v.token.String(),
		},
		InsecureSkipTLS: true,
	})
	if err != nil {
		return nil, err
	}
	return repository, nil
}

func (v Vcs) PushRepositoryToProject(project *domain.Project, repository *git.Repository) error {
	_, err := repository.CreateRemote(&config.RemoteConfig{
		Name:  defaultRemoteName,
		URLs:  []string{project.VcsUrl},
		Fetch: []config.RefSpec{"+refs/heads/*:refs/remotes/origin/*"},
	})
	if err != nil {
		return err
	}

	err = repository.Push(&git.PushOptions{
		RemoteName: defaultRemoteName,
		Auth: &githttp.BasicAuth{
			Username: "nobody",
			Password: v.token.String(),
		},
		InsecureSkipTLS: true,
	})
	if err != nil {
		return err
	}

	return nil
}
