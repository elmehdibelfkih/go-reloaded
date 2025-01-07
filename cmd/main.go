package main

import (
	"fmt"
	"go-reloaded/internal"
	"os"
)

func main() {
	filesName := os.Args[1:]
	if len(filesName) != 2 {
		fmt.Fprintf(os.Stderr, "Error: program args must be 2\n")
		fmt.Fprintf(os.Stderr, "Usage: go run cmd/. [input] [output]\n")
		return
	}
	input, err := os.Open(filesName[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}
	defer input.Close()
	if filesName[0] == filesName[1] {
		fmt.Fprintf(os.Stderr, "Error: the input must be different from the output\n")
		return
	}
	output, err := os.Create(filesName[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}
	defer output.Close()
	internal.HandelLine(input, output)
}
