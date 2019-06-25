package main

import (
	"fmt"
	"net/http"
)

const (
	route string = "/"
	port  string = ":8079"
)

func printer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "it is a good day to die")
}

func main() {
	http.HandleFunc(route, printer)
	fmt.Println("сервер слушает порт", port)
	http.ListenAndServe(port, nil)
}
