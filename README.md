### About

Simple test of few RAML Go parsers. Tests simply try to parse a set of examples and report if parser returned an error.

A fine collection of RAML files can be composed each containing one/few RAML features to test those in isolation.

### Install

```sh
git clone git@github.com:postatum/raml-parsers-test.git
cd raml-parsers-test
```

*Optional:*
Install package with `go install .` to run it as a binary `raml-parsers-test` instead of `go run *.go`.

### Run

```sh
go run *.go -parser jumpscale
go run *.go -parser goraml
go run *.go -parser tsaikd
```

### Other options

Help:

```sh
go run *.go -h
```

Verbose output (prints errors):

```sh
go run *.go -parser PARSER_NAME -verbose
```

Using different examples folder:

```sh
go run *.go -parser PARSER_NAME -examples ../some/folder/with_raml

```