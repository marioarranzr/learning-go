package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/golang-collections/collections/queue"
	"github.com/julienschmidt/httprouter"
)

type interfaceList []interface{}

var tpl *template.Template
var q *queue.Queue

// list to check the result is correct
var l interfaceList

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	q = queue.New()
	q.Enqueue(2)
	q.Enqueue(6)
	q.Enqueue(4)
	q.Enqueue(1)
	q.Enqueue(3)
	l = append(l, 2, 6, 4, 1, 3)
}

func main() {

	mux := httprouter.New()
	mux.GET("/", index)
	mux.GET("/enqueue", enqueue)
	mux.POST("/enqueue", postEnqueue)
	mux.GET("/dequeueAll", dequeueAll)
	mux.GET("/dequeue", dequeue)

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
	m := make(map[string]interfaceList)
	m["List"] = l
	err := tpl.ExecuteTemplate(w, "index.gohtml", m)
	HandleError(w, err)
}

func dequeueAll(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var values interfaceList
	for q.Len() != 0 {
		data := q.Dequeue()
		values = append(values, data)
	}
	l = make(interfaceList, 0)

	m := make(map[string]interfaceList)
	m["Queue"] = values
	m["List"] = l
	err := tpl.ExecuteTemplate(w, "index.gohtml", m)
	HandleError(w, err)
}

func dequeue(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var values interfaceList
	data := q.Dequeue()
	values = append(values, data)
	l = l[1:]

	m := make(map[string]interfaceList)
	m["Queue"] = values
	m["List"] = l
	err := tpl.ExecuteTemplate(w, "index.gohtml", m)
	HandleError(w, err)
}

func enqueue(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "queue-enqueue.gohtml", nil)
	HandleError(w, err)
}

func postEnqueue(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	value := rand.Intn(100)
	q.Enqueue(value)

	var values interfaceList
	values = append(values, value)
	l = append(l, value)

	m := make(map[string]interfaceList)
	m["Queue"] = values
	m["List"] = l
	err := tpl.ExecuteTemplate(w, "index.gohtml", m)
	HandleError(w, err)
}

// HandleError Error response haldling
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
