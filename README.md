# erago
Erajaya CLI generate project.

## Installation

You can install by using go binary.

```shell
go install github.com/spacetronot-research-team/erago@latest
```

or you can define your prefered version.

```shell
go install github.com/spacetronot-research-team/erago@v0.0.5
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
