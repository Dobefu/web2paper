# Web2Paper

[![CI Status](https://github.com/Dobefu/web2paper/actions/workflows/ci.yml/badge.svg)](https://github.com/Dobefu/web2paper/actions/workflows/ci.yml)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Dobefu_web2paper&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=Dobefu_web2paper)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=Dobefu_web2paper&metric=coverage)](https://sonarcloud.io/summary/new_code?id=Dobefu_web2paper)
[![Go Report Card](https://goreportcard.com/badge/github.com/Dobefu/web2paper)](https://goreportcard.com/report/github.com/Dobefu/web2paper)

## Installation

To get started, install the latest version of the application:

```sh
go install github.com/Dobefu/web2paper@latest
```

Or, if you prefer building from source:

```sh
git clone https://github.com/Dobefu/web2paper.git
cd web2paper
go build
```

## Usage

Convert HTML to PDF with a single command:

```sh
./web2paper convert -i input.html -o output.pdf
```

- `-i` or `--input`: Path to the input HTML file (required)
- `-o` or `--output`: Path for the output PDF (required)
