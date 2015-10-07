package main

import (
	"flag"
	"fmt"
	"os"
)

const Version = "0.0.1"

var version = flag.Bool("version", false, "print version and exit")

func bold(str string) string {
	return "\033[1m" + str + "\033[0m"
}

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if *version {
		fmt.Printf("utmptail %s\n", Version)
		os.Exit(0)
	}

}
