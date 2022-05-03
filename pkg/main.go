package main

import (
	_ "embed"
	"flag"
	"io/ioutil"

	"github.com/alecthomas/repr"
)

func main() {

	path := flag.String("in", "", "path to FIDL file to parse")
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
