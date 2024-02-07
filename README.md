# erago
Erajaya CLI generate project.

![erago](https://github.com/spacetronot-research-team/erago/assets/57469556/10dc6e4c-25e7-4b48-bb9e-1e34348b8012)

```
├── cmd/                Initial stage of the application will run.
├── internal/           Core module of the application and contains the implementation of various business logic.
│   ├──  controller/    This module is only to gather input (REST/gRPC/console/etc) and pass input as request to service.
│   │   ├──  http/
│   │   ├──  grpc/
│   ├──  service/       This module contain business logic, this module get input request from controller, this module use repository for things related to persistence.
│   ├──  repository/    This module only for things related to persistence (CRUD database/redis/etc).
├── go.mod
└── go.sum
```

## Installation

You can install by using go binary.

```shell
go install github.com/spacetronot-research-team/erago@latest
```

or you can define your prefered version.

```shell
go install github.com/spacetronot-research-team/erago@v0.0.19
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