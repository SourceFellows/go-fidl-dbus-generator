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
	outFile := flag.String("out", "", "path to generated file")

	packageName := flag.String("package", "", "package to generate the result in")

	var writerType pkg.WriterType
	generateReceiver := flag.Bool("receiver", false, "generate receiver impl")
	generateSender := flag.Bool("sender", false, "generate sender impl")

	debug := flag.Bool("debug", false, "debug mode")

	flag.Parse()

	if inFile == nil || *inFile == "" {
		log.Println("no input file given")
		flag.PrintDefaults()
		return
	}

	if (generateReceiver != nil && generateSender != nil) && (!*generateReceiver && !*generateSender) {
		log.Println("you should decide if you want a receiver or sender impl")
		flag.PrintDefaults()
		return
	}

	if *generateReceiver && *generateSender {
		log.Println("you can generate receiver OR sender impl")
		flag.PrintDefaults()
		return
	}

	if *generateSender {
		writerType = pkg.SenderWriter
	} else {
		writerType = pkg.ReceiverWriter
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
		fidl.TargetPackage = *packageName
	} else {
		fidl.TargetPackage = fidl.PackageInfo.Name
	}

	out := os.Stdout
	if outFile != nil && *outFile != "" {
		out, err = os.OpenFile(*outFile, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
	}

	err = pkg.Write(fidl, writerType, out)
	if err != nil {
		log.Fatalln(err)
	}
}
