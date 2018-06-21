package main

import (
    raml "github.com/Jumpscale/go-raml/raml"
)

type Parser func(string) error

func Jumpscale(fpath string) error {
    apiDef := &raml.APIDefinition{}
    return raml.ParseFile(fpath, apiDef)
}
