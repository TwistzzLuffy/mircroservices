package Handlers

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
	return &Engine{}
}

func (e *Engine) ServerHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", d)
}
