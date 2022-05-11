package main

import (
	_ "embed"
	"flag"
	"github.com/SourceFellows/go-fidl-dbus-generator/pkg"
	"github.com/alecthomas/repr"
	"log"
	"os"
)

func main() {

	inFile := flag.String("in", "", "path to FIDL file to parse")
	flag.Parse()

	if inFile == nil || *inFile == "" {
		log.Println("no input file given")
		flag.PrintDefaults()
		return
	}

	file, err := os.Open(*inFile)
	if err != nil {
		log.Fatalf("error while reading in file: %v", err)
	}
	defer file.Close()

	parser := pkg.NewParser(file)

	fidl, err := parser.Parse()
	if err != nil {
		log.Fatalln(err)
	}

	repr.Println(fidl)

	err = pkg.Write(fidl, os.Stdout)
	if err != nil {
		log.Fatalln(err)
	}
}
