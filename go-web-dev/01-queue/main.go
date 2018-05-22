package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	_ "go-web-dev/01-queue/queue"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	mux := httprouter.New()
	mux.GET("/", index)
	mux.GET("/about", about)
	mux.GET("/queue", queue)

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Minute,
		Handler:      mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func queue(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// items := []int{1, 2, 3}

	s := ItemQueue{}
	s.Enqueue(1)
	s.Enqueue(2)
	s.Enqueue(3)

	// if size := s.Size(); size != 3 {
	// 	fmt.Println("dddd")
	// }
}

// HandleError Error response haldling
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
