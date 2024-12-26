package main

import (
	"fmt"
	"os"
	"go-reloaded/internal"
)

func main() {
	filesName := os.Args[1:]
	input, err := os.Open(filesName[0])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer input.Close()
	output, err := os.Create(filesName[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer output.Close()
	internal.HandelLine(input, output)
}