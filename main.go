package main

import (
	"flag"
	"fmt"
)

func main() {
	parserFl := flag.String(
		"parser", "jumpscale",
		"Parser to test. Supported: jumpscale, goraml, tsaikd.")
	examplesFl := flag.String(
		"examples", "./raml_examples",
		"Examples folder path")
	verboseFl := flag.Bool(
		"verbose", false,
		"Verbose mode. Parsing errors will be printed.")
	flag.Parse()

	verbose := *verboseFl
	parsers := map[string]Parser{
		"jumpscale": Jumpscale,
		"goraml":    Goraml,
		"tsaikd":    Tsaikd,
	}
	parser, ok := parsers[*parserFl]
	if !ok {
		fmt.Println("Not supported parser. See help (-h).")
		return
	}

	fileList, err := ListRamls(*examplesFl)
	if err != nil {
		fmt.Printf("Failed to list RAML files: %s\n", err)
		return
	}

	var result string
	for _, fpath := range fileList {
		result = "OK"
		err = parser(fpath)
		if err != nil {
			result = "FAIL"
		}
		fmt.Printf("> %s: %s\n", fpath, result)
		if verbose && err != nil {
			fmt.Println(err.Error())
		}
	}
}
