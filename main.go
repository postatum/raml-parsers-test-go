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
		"examples", "",
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

	examples := *examplesFl
	if examples == "" {
		fmt.Println("Missing examples path param")
		return
	}

	fileList, err := ListRamls(examples)
	if err != nil {
		fmt.Printf("Failed to list RAML files: %s\n", err)
		return
	}

	var result string
	for _, fpath := range fileList {
		fmt.Printf("> Parsing %s: ", fpath)
		err, notPanic := parser(fpath)
		if !notPanic {
			continue
		}
		result = "OK"
		if err != nil {
			result = "FAIL"
		}
		fmt.Printf("%s\n", result)
		if verbose && err != nil {
			fmt.Println(err.Error())
		}
	}
}
