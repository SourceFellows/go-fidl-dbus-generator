package main

import (
	_ "embed"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/SourceFellows/go-fidl-dbus-generator/pkg"
	gofidl "github.com/SourceFellows/go-fidl-dbus-generator/pkg"
	"github.com/alecthomas/repr"
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

	fidl, err := gofidl.ParseFidl(file)
	if err != nil {
		log.Fatalln(err)
	}

	repr.Println(fidl)

	fmt.Println(pkg.Write(fidl, os.Stdout))

}
