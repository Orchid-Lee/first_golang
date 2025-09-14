package main

import (
	_ "first_golang/lib2"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Server start...")
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe("localhost:1234", nil); err != nil {
		log.Fatal("ListenAndServe err: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	s := "hello world"
	fmt.Fprintf(w, "%s", s)
	log.Printf("%s", s)
}
