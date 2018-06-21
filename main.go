package main

import (
    "fmt"
    "flag"
)

func main()  {
    parserFl := flag.String(
        "parser", "jumpscale",
        "Parser to test. Supported: jumpscale, goraml, tsaikd.")
    verboseFl := flag.Bool(
        "verbose", false,
        "Verbose mode. Parsing errors will be printed.")
    flag.Parse()

    parsers := map[string]Parser{
        "jumpscale": Jumpscale,
        // ADD MORE
    }
    parser, ok := parsers[*parserFl]
    if !ok {
        fmt.Println("Not supported parser. See help (-h).")
        return
    }

    fileList, err := ListRamls("./raml_examples")
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
        fmt.Printf("%s: %s\n", fpath, result)
        if *verboseFl {
            fmt.Println(err.Error())
        }
    }
}
