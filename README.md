# erago
Erajaya CLI generate project.

![erago](https://github.com/spacetronot-research-team/erago/assets/57469556/10dc6e4c-25e7-4b48-bb9e-1e34348b8012)

```
├── cmd/
│   ├── main.go             Initial stage of the application will run.
├── database/
│   ├── migrate/
│   │   ├── up.go           Database migrate up, 'go run database/migrate/up.go'.
│   ├── schema_migration/   Contain database migrate sql file.
├── docs/
│   ├── errors.json         Contain all errors list to be deplayed by frontend.
├── internal/
│   ├── controller/         Contain things related to gather input (REST/gRPC/console/etc) and pass input as request to service.
│   │   ├── http/
│   │   ├── grpc/
│   ├── repository/         Contain things related to persistence (CRUD database/redis/etc).
│   │   ├── mockrepository/
│   ├── router/
│   │   ├── injection.go    Contain dependency injection from controller to service to repository.
│   ├── service/            Contain business logic, this package get input request from controller, this package use repository for things related to persistence.
│   │   ├── mockservice/
├── go.mod
└── go.sum
└── README.md
```

## Installation

You can install by using go binary.

```shell
go install github.com/spacetronot-research-team/erago@latest
```

or you can define your prefered version.

```shell
go install github.com/spacetronot-research-team/erago@v0.0.24
```

Or you can download erago binary from [release page](https://github.com/spacetronot-research-team/erago/releases).

You can check your version by running.

```shell
erago version
```

## Get Started

1. Follow installation.

2. Create new project.

```shell
erago create-project go-customer github.com/spacetronot-research-team/go-customer
```

New project will be created in directory `go-customer`.

3. Create new domain.

```shell
cd go-customer && erago create-domain profile
```

New domain will be created in directory `internal/controller/http/` and `internal/service/` and `internal/repository/`.

## Docs

```shell
hidayat@thinkubuntu:~$ erago --help
Erajaya CLI generate project.

Usage:
  erago [command]

Available Commands:
  create-domain  Create new domain with the provided domain name
  create-project Create new project with the provided domain name
  explain        Explain code architecture
  help           Help about any command
  version        Print erago version

Flags:
  -h, --help   help for erago

Use "erago [command] --help" for more information about a command.
```