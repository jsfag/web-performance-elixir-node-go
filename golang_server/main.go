package main

import (
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// message := r.URL.Path
	// message = strings.TrimPrefix(message, "/")
	// message = "Hello " + message
	w.Write([]byte("world"))
}
func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
