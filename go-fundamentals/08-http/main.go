package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com") //response, error
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, resp.Body) // Print by console (this Writter could be also a file) all the data in resp.Body
	// io.Copy(os.Stdout, resp.Body) // Print by console (this Writter could be also a file) all the data in resp.Body

	// bs := make([]byte, 999999)
	// resp.Body.Read(bs) // Read the data into resp.Body and copy it into the []byte
	// fmt.Println(string(bs))
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
