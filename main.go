package main

import (
	"net/http"
)

func handle(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/", handle)
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
