### TODO:

* [x] Write test suite against built-in jumpscale examples
* [x] Write test-suite for jumpscale/go-raml/raml
* [ ] Write test-suite for go-raml/raml
* [ ] Write test-suite for tsaikd/go-raml-parser
* [ ] Replace jumpscale examples with raml-org/raml-examples


### About

Simple test of few RAML Go parsers. Tests simply try to parse a set of examples and report if parser returned an error.

A fine collection of RAML files can be composed each containing one/few RAML features to test those in isolation.

### Run

```sh
go run *.go -parser jumpscale
go run *.go -parser goraml
go run *.go -parser tsaikd
```

### Verbose output (prints errors)

```sh
go run *.go -parser PARSER_NAME -verbose
```

### Help

```sh
go run *.go -h
```