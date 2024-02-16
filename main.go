package main

// TODO:
// - use something like go-flags which supports required values?
// 		+ https://github.com/jessevdk/go-flags

import (
	"flag"

	"github.com/btoll/validator/validators"
)

func main() {
	file := flag.String("file", "deployment.json", "The name of the file to validate")
	flag.Parse()

	validators.New(*file)
}
