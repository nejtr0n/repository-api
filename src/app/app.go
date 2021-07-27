package app

import "github.com/nejtr0n/repository-api/src/app/command"

func NewApplication(createProjectFromDull command.CreateProjectFromDullHandler) Application {
	return Application{
		Commands: Commands{
			CreateProjectFromDull: createProjectFromDull,
		},
	}
}

type Application struct {
	Commands Commands
}

type Commands struct {
	CreateProjectFromDull command.CreateProjectFromDullHandler
}
