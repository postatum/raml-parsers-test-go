package main

import (
	raml "github.com/Jumpscale/go-raml/raml"
)

// Parser is a parsing type function
type Parser func(string) error

// Jumpscale runs jumpscale/go-raml/raml parser
func Jumpscale(fpath string) error {
	apiDef := &raml.APIDefinition{}
	return raml.ParseFile(fpath, apiDef)
}
