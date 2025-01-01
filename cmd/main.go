package main

import (
	"fmt"
	"go-reloaded/internal"
	"os"
)

func main() {
	filesName := os.Args[1:]
	if len(filesName) != 2 {
		fmt.Println("Error: program args must be 2")
		return
	}
	input, err := os.Open(filesName[0])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer input.Close()
	if filesName[0] == filesName[1] {
		fmt.Println("Error: the input must be different from the output")
		return
	}
	output, err := os.Create(filesName[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer output.Close()
	internal.HandelLine(input, output)
}
