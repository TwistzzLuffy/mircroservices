package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Engine struct {
	l *log.Logger
}

func NewEngine(l *log.Logger) *Engine {
	return &Engine{l}
}

func (e *Engine) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	e.l.Println("Hello World")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", d)
}
