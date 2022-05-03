package main

import (
	_ "embed"
	"flag"
	"github.com/alecthomas/repr"
	"io/ioutil"
)

func main() {
	path := flag.String("path", "", "path to FIDL file to parse")
	flag.Parse()

	file, err := ioutil.ReadFile(*path)
	if err != nil {
		panic(err)
	}

	fidl, err := parseFidl(file)
	if err != nil {
		panic(err)
	}

	repr.Println(fidl)

}
