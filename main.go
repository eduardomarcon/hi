package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hi")
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("server running")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
