// start server and client service by execute both client.exe and server.exe
// test via url, connect: http://localhost:8080/add/2/2
package main

import (
	"encoding/json"
	"github.com/marc47marc47/grpc-protobuf/proto"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type Message struct {
	Message string `json:"message"`
	Code    int    `json:"status"`
}

var client proto.AddServiceClient

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client = proto.NewAddServiceClient(conn)
	if client != nil {
		log.Println("client connected")
	}
	router := mux.NewRouter()
	router.HandleFunc("/", Welcome).Methods("GET")
	router.HandleFunc("/add/{a}/{b}", Sum).Methods("GET")
	router.HandleFunc("/multiply/{a}/{b}", Multiply).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	var m = &Message{"welcome to grpc server test", http.StatusOK}
	WithError(w, *m)
}
func Sum(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	a, err := strconv.ParseInt(params["a"], 10, 64)
	if err != nil {
		panic(err)
	}
	b, err := strconv.ParseInt(params["b"], 10, 64)
	if err != nil {
		panic(err)
	}
	req := &proto.Request{A: int64(a), B: int64(b)}
	if response, err := client.Add(r.Context(), req); err == nil {
		var m = &Message{strconv.FormatInt(response.Result, 10), http.StatusOK}
		WithError(w, *m)
	} else {
		var m = &Message{"some bad request", http.StatusBadRequest}
		WithError(w, *m)
	}
}
func Multiply(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	a, err := strconv.ParseInt(params["a"], 10, 64)
	if err != nil {
		panic(err)
	}
	b, err := strconv.ParseInt(params["b"], 10, 64)
	if err != nil {
		panic(err)
	}
	req := &proto.Request{A: int64(a), B: int64(b)}
	if response, err := client.Multiply(r.Context(), req); err == nil {
		var m = &Message{strconv.FormatInt(response.Result, 10), http.StatusOK}
		WithError(w, *m)
	} else {
		var m = &Message{"some bad request", http.StatusBadRequest}
		WithError(w, *m)
	}
}

func Json(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		WithError(w, Message{"invalid request", http.StatusBadRequest})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func WithError(w http.ResponseWriter, m Message) {
	response, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(m.Code)
	w.Write(response)
}
