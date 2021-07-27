//+build wireinject

package main

import (
	"github.com/nejtr0n/repository-api/src/adapters/git"
	"github.com/nejtr0n/repository-api/src/adapters/gitlab"
	"github.com/nejtr0n/repository-api/src/app"
	"github.com/nejtr0n/repository-api/src/app/command"
	"github.com/nejtr0n/repository-api/src/domain"
	"github.com/nejtr0n/repository-api/src/ports"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

func InitRouter(config *cli.Context) (*gin.Engine, error) {
	wire.Build(
		ports.NewRouter,
		ports.NewHttpServer,
		app.NewApplication,
		command.NewCreateProjectFromDullHandler,
		domain.NewGitlabApiToken,
		domain.NewGitlabBaseUrl,
		domain.NewService,
		domain.NewProviderFactory,
		gitlab.NewProvider,
		git.NewVcs,
		wire.Bind(new(domain.GitlabProvider), new(*gitlab.Provider)),
		wire.Bind(new(domain.GitlabVcs), new(*git.Vcs)),
	)

	return &gin.Engine{}, nil
}
