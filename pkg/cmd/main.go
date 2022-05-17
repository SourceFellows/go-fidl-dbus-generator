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

	packageName := flag.String("package", "", "package to generate the result in")

	var writerType pkg.WriterType
	generateServer := flag.Bool("server", false, "generate server impl")
	generateClient := flag.Bool("client", false, "generate client impl")

	debug := flag.Bool("debug", false, "debug mode")

	flag.Parse()

	if inFile == nil || *inFile == "" {
		log.Println("no input file given")
		flag.PrintDefaults()
		return
	}

	if (generateServer != nil && generateClient != nil) && (!*generateServer && !*generateClient) {
		log.Println("you should decide if you want a server or client impl")
		flag.PrintDefaults()
		return
	}

	if *generateServer && *generateClient {
		log.Println("you can generate server OR client impl")
		flag.PrintDefaults()
		return
	}

	if *generateClient {
		writerType = pkg.ClientWriter
	} else {
		writerType = pkg.ServerWriter
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

	if debug != nil && *debug {
		repr.Println(fidl)
	}

	if *packageName != "" {
		fidl.PackageInfo.Name = *packageName
	}

	err = pkg.Write(fidl, writerType, os.Stdout)
	if err != nil {
		log.Fatalln(err)
	}
}
