package router

import (
	"github.com/spacetronot-research-team/erago/internal/controller/cli"
	"github.com/spacetronot-research-team/erago/internal/repository"
	"github.com/spacetronot-research-team/erago/internal/service"
)

func getVersionController() *cli.VersionController {
	versionService := service.NewVersionService()
	versionController := cli.NewVersionController(versionService)
	return versionController
}

func getExplainController() *cli.ExplainController {
	explainRepository := repository.NewExplainRepository()
	explainService := service.NewExplainService(explainRepository)
	explainController := cli.NewExplainController(explainService)
	return explainController
}

func GetDomainController() *cli.DomainController {
	domainRepository := repository.NewDomainRepository()
	domainService := service.NewDomainService(domainRepository)
	domainController := cli.NewDomainController(domainService)
	return domainController
}

func getProjectController() *cli.ProjectController {
	domainRepository := repository.NewDomainRepository()
	domainService := service.NewDomainService(domainRepository)

	projectRepository := repository.NewProjectRepository()
	projectService := service.NewProjectService(projectRepository, domainService)
	projectController := cli.NewProjectController(projectService)
	return projectController
}
