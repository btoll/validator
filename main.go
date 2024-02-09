package main

// TODO:
// - use something like go-flags which supports required values?
// 		+ https://github.com/jessevdk/go-flags

import (
	"flag"

	"github.com/btoll/validator/validators"
)

func main() {
	file1 := flag.String("file1", "deployment.json", "The name of the file to validate")
	file2 := flag.String("file2", "deployment.json", "The name of the other file to validate")
	//	raw := flag.Bool("raw", false, "Print the raw unmarshaled JSON")
	flag.Parse()

	v := validators.New(
		validators.NewDocument(*file1),
		validators.NewDocument(*file2))
	v.Validate()
}
