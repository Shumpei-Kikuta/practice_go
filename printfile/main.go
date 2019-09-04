package main

import (
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println(string(data[:count]))
}
