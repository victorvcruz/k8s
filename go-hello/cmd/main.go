package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	fmt.Println("Server started at port 9090")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 9090), nil))
}
