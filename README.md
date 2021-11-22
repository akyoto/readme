# readme

[![Godoc][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][coverage-image]][coverage-url]
[![Sponsor][sponsor-image]][sponsor-url]

Readme generator.

## Installation

```shell
go get -u github.com/akyoto/readme/...
```

## Usage

```shell
readme
```

Generates `README.md` files for every subdirectory that has a `README.src.md` file.

## Snippets

```text
name
```

Repository name.

```text
go:header
```

Header for Go projects.

```text
go:footer
```

Footer for Go projects.

```text
go:install
```

Installation instructions for Go projects.

[godoc-image]: https://godoc.org/github.com/akyoto/readme?status.svg
[godoc-url]: https://godoc.org/github.com/akyoto/readme
[report-image]: https://goreportcard.com/badge/github.com/akyoto/readme
[report-url]: https://goreportcard.com/report/github.com/akyoto/readme
[tests-image]: https://cloud.drone.io/api/badges/akyoto/readme/status.svg
[tests-url]: https://cloud.drone.io/akyoto/readme
[coverage-image]: https://codecov.io/gh/akyoto/readme/graph/badge.svg
[coverage-url]: https://codecov.io/gh/akyoto/readme
[sponsor-image]: https://img.shields.io/badge/github-donate-green.svg
[sponsor-url]: https://github.com/users/akyoto/sponsorship
