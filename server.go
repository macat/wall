package main

import (
	"github.com/garyburd/redigo/redis"
	"io"
	"log"
	"net/http"
	"os"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	conn, err := redis.DialURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	server := http.Server{
		Addr:    ":8000",
		Handler: &wallHandler{},
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/canvases"] = canvases

	server.ListenAndServe()

}

func canvases(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

type wallHandler struct{}

func (*wallHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	io.WriteString(w, "My server: "+r.URL.String())
}
