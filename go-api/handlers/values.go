package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Value struct {
	l *log.Logger
}

func NewValue(l *log.Logger) *Value {
	rand.Seed(42)
	return &Value{l}
}

func (value Value) GetValues(w http.ResponseWriter, r *http.Request) {
	x := rand.Intn(10)
	value.l.Printf("Handle Get requests -> Duration:%d", x)
	time.Sleep(time.Duration(x) * time.Second)
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(map[string]string{"result": "success", "duration": fmt.Sprintf("%d seconds", x)})
	w.Write(response)
}

func (value Value) PutValue(w http.ResponseWriter, r *http.Request) {
	x := rand.Intn(10)
	value.l.Printf("Handle Put requests -> Duration:%d", x)
	time.Sleep(time.Duration(x) * time.Second)
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(map[string]string{"result": "success", "duration": fmt.Sprintf("%d seconds", x)})
	w.Write(response)
}

func (Value Value) PostValue(w http.ResponseWriter, r *http.Request) {
	x := rand.Intn(10)
	Value.l.Printf("Handle Post requests -> Duration:%d", x)
	time.Sleep(time.Duration(x) * time.Second)
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(map[string]string{"result": "success", "duration": fmt.Sprintf("%d seconds", x)})
	w.Write(response)
}
