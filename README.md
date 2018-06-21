# About

Simple test of few RAML Go parsers. Tests simply try to parse a set of examples and report if parser returned an error.

A fine collection of RAML files can be composed each containing one/few RAML features to test those in isolation.

# Install & run as bin

```sh
go install github.com/postatum/raml-parsers-test
git clone git@github.com:raml-org/raml-examples.git raml_examples
raml-parsers-test -parser PARSER_NAME -examples ./raml_examples
```

# Install & run raw code

```sh
git clone git@github.com:postatum/raml-parsers-test.git
git clone git@github.com:raml-org/raml-examples.git raml_examples
cd raml-parsers-test
go run *.go -parser PARSER_NAME -examples ../raml_examples

```

# Options

Help:

```sh
raml-parsers-test -h
```

Parser (defaults to `jumpscale`):
```sh
raml-parsers-test -parser jumpscale/goraml/tsaikd
```

Verbose output (prints errors) (defaults to `false`):

```sh
raml-parsers-test -parser PARSER_NAME
```

Examples folder (**required**):

```sh
raml-parsers-test -parser PARSER_NAME -examples ../some/folder/with_raml

```
