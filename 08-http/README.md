## Simple HTTP server

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	// http.Handle("/", http.HandlerFunc(handler))
	// http.Handle("/", http.TimeoutHandler(http.HandlerFunc(handler), time.Second, "timed out"))


	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
}
```

## Simple HTTP client

```go
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "http://localhost:8080"

	cli := &http.Client{
		Timeout: time.Second,
	}

	res, err := cli.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("successfully connected to", url)
}
```

>`go run main.go`
