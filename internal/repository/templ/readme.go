package templ

import "github.com/spacetronot-research-team/erago/pkg/version"

//nolint:lll
func GetReadmeTemplate() string {
	var readmeTemplate string

	readmeTemplate += "# {{.ProjectName}}\n"
	readmeTemplate += "Generated by erago@" + version.Current + ".\n"
	readmeTemplate += "\n"
	readmeTemplate += "![code-arch-explain](https://github.com/spacetronot-research-team/erago/assets/57469556/1b407531-acd4-4d6b-89e8-47175b8482a6)\n"
	readmeTemplate += "\n"
	readmeTemplate += "```\n"
	readmeTemplate += "├── cmd/\n"
	readmeTemplate += "│   ├── main.go             Initial stage of the application will run.\n"
	readmeTemplate += "├── database/\n"
	readmeTemplate += "│   ├── migrate/\n"
	readmeTemplate += "│   │   ├── up.go           Database migrate up, 'go run database/migrate/up.go'.\n"
	readmeTemplate += "│   ├── schema_migration/   Contain database migrate sql file.\n"
	readmeTemplate += "├── docs/\n"
	readmeTemplate += "│   ├── errors.json         Contain all errors list to be deplayed by frontend.\n"
	readmeTemplate += "├── internal/\n"
	readmeTemplate += "│   ├── controller/         Contain things related to gather input (REST/gRPC/console/etc) and pass input as request to service.\n"
	readmeTemplate += "│   │   ├── http/\n"
	readmeTemplate += "│   │   ├── grpc/\n"
	readmeTemplate += "│   ├── repository/         Contain things related to persistence (CRUD database/redis/etc).\n"
	readmeTemplate += "│   ├── router/\n"
	readmeTemplate += "│   │   ├── injection.go    Contain dependency injection from controller to service to repository.\n"
	readmeTemplate += "│   ├── service/            Contain business logic, this package get input request from controller, this package use repository for things related to persistence.\n"
	readmeTemplate += "├── go.mod\n"
	readmeTemplate += "└── go.sum\n"
	readmeTemplate += "└── README.md\n"
	readmeTemplate += "```\n"

	return readmeTemplate
}
