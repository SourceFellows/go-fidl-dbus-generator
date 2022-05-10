package main

import (
	"bytes"
	_ "embed"
	"flag"
	"github.com/SourceFellows/go-fidl-dbus-generator/pkg/lexer"
	"github.com/alecthomas/repr"
	"io/ioutil"
	"log"
)

func main() {

	inFile := flag.String("in", "", "path to FIDL file to parse")
	flag.Parse()

	if inFile == nil || *inFile == "" {
		log.Println("no input file given")
		flag.PrintDefaults()
		return
	}

	file, err := ioutil.ReadFile(*inFile)
	if err != nil {
		log.Fatalf("error while reading in file: %v", err)
	}

	parser := lexer.NewParser(bytes.NewReader(file))

	fidl, err := parser.Parse()
	if err != nil {
		log.Fatalln(err)
	}

	repr.Println(fidl)
}
