# erago
Erajaya CLI generate project.

![erago](https://github.com/spacetronot-research-team/erago/assets/57469556/8b25595e-1e07-4605-bb25-7b26e335c711)

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
go install github.com/spacetronot-research-team/erago@v0.0.6
```

Or you can download erago binary from release page.

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

New domain will be created in directory `go-customer/internal`.
