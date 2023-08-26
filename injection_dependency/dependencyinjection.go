package main

import (
	"fmt"
	"io"
	"net/http"
)

func Hello(escritor io.Writer, nome string) {
	fmt.Fprintf(escritor, "Hello, %s", nome)
}

func HandlerMyHello(w http.ResponseWriter, r *http.Request) {
	Hello(w, "World")
}

func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(HandlerMyHello))
	if err != nil {
		fmt.Println(err)
	}
}
