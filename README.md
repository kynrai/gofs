# ttz (10100-cli)

`ttz` is a 10100 CLI tool to generate opinionated secure projects

## Installation

```bash
git clone git@github.com:atos-digital/ttz.git
cd ttz
go mod tidy
go install
```

## Usage

```bash
ttz
```

## Using generated templates

The template includes several modules that are optional and should be deleted to reduce build size. For example we include a postgres connector and a cloudsql connector for conveniance, but you should likely only need one of them.
