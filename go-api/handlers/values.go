package handlers

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Value struct {
	l *log.Logger
}

func NewValue(l *log.Logger) {
	return &Value{l}
}

func (value Value) GetValues(w http.ResponseWriter, r *http.Request) {
	x := rand.Int(10)
	value.l.Printf("Handle Get requests -> Duration:%d", x)
	time.Sleep(x * time.Second)
	w.WriteHeader(http.StatusOK)
}

func (value Value) PutValues(w http.ResponseWriter, r *http.Request) {
	x := rand.Int(10)
	value.l.Printf("Handle Put requests -> Duration:%d", x)
	time.Sleep(x * time.Second)
	w.WriteHeader(http.StatusOK)
}

func (Value Value) PostValues(w http.ResponseWriter, r *http.Request) {
	x := rand.Int(10)
	Value.l.Printf("Handle Post requests -> Duration:%d", x)
	time.Sleep(x * time.Second)
	w.WriteHeader(http.StatusOK)
}
