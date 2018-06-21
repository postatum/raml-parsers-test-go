package main

import (
	jumpscale "github.com/Jumpscale/go-raml/raml"
	tsaikd "github.com/tsaikd/go-raml-parser/parser"
	goraml "gopkg.in/raml.v0" // github.com/go-raml/raml
)

// Parser is a parsing type function
type Parser func(string) error

// Jumpscale runs jumpscale/go-raml/raml parser
func Jumpscale(fpath string) error {
	apiDef := &jumpscale.APIDefinition{}
	return jumpscale.ParseFile(fpath, apiDef)
}

// Goraml runs go-raml/raml parser
func Goraml(fpath string) error {
	_, err := goraml.ParseFile(fpath)
	return err
}

// Tsaikd runs tsaikd/go-raml-parser/parser parser
func Tsaikd(fpath string) error {
	ramlParser := tsaikd.NewParser()
	_, err := ramlParser.ParseFile(fpath)
	return err
}
