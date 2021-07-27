package domain

import "github.com/go-git/go-git/v5"

type Vcs interface {
	CloneRepository(project *Project) (*git.Repository, error)
	PushRepositoryToProject(project *Project, repository *git.Repository) error
}

type GitlabVcs interface {
	Vcs
}
