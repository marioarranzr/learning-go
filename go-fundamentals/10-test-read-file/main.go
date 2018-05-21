package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	data := make([]byte, 1000000)
	count, err := file.Read(data)

	if err != nil {
		log.Fatal(err)
	}

	if count > 0 {
		fmt.Println(string(data))
	}

	return

}
