package main

import (
	"fmt"
	"os"
	"go-reloaded/internal"
)

func main() {
	filesName := os.Args[1:]
	if len(filesName) != 2 {
		fmt.Println("Error: to many args")
		return
	}
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