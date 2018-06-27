package main

import (
	"flag"
	"fmt"
)

func main() {
	parserFl := flag.String(
		"parser", "jumpscale",
		"Parser to test. Supported: jumpscale, goraml, tsaikd.")
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

	examplesFl := CloneTckRepo()
	fileList, err := ListRamls(examplesFl)
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
		failed := err != nil
		if ShouldFail(fpath) {
			failed = !failed
		}
		result = "OK"
		if failed {
			result = "FAIL"
		}
		fmt.Printf("%s\n", result)
		if verbose && err != nil {
			fmt.Println(err.Error())
		}
	}
}
