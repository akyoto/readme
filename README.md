# readme

[![Godoc][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][coverage-image]][coverage-url]
[![Patreon][patreon-image]][patreon-url]

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

## Style

Please take a look at the [style guidelines](https://github.com/akyoto/quality/blob/master/STYLE.md) if you'd like to make a pull request.

## Sponsors

| [![Scott Rayapoullé](https://avatars3.githubusercontent.com/u/11772084?s=70&v=4)](https://github.com/soulcramer) | [![Eduard Urbach](https://avatars2.githubusercontent.com/u/438936?s=70&v=4)](https://twitter.com/eduardurbach) |
| --- | --- |
| [Scott Rayapoullé](https://github.com/soulcramer) | [Eduard Urbach](https://eduardurbach.com) |

Want to see [your own name here](https://www.patreon.com/eduardurbach)?

[godoc-image]: https://godoc.org/github.com/akyoto/readme?status.svg
[godoc-url]: https://godoc.org/github.com/akyoto/readme
[report-image]: https://goreportcard.com/badge/github.com/akyoto/readme
[report-url]: https://goreportcard.com/report/github.com/akyoto/readme
[tests-image]: https://cloud.drone.io/api/badges/akyoto/readme/status.svg
[tests-url]: https://cloud.drone.io/akyoto/readme
[coverage-image]: https://codecov.io/gh/akyoto/readme/graph/badge.svg
[coverage-url]: https://codecov.io/gh/akyoto/readme
[patreon-image]: https://img.shields.io/badge/patreon-donate-green.svg
[patreon-url]: https://www.patreon.com/eduardurbach
