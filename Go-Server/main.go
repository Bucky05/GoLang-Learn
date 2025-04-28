package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8000\n")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(resp, "404 not found", http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(resp, "Method not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(resp, "Hello ther")
	//fmt.Fprintf(resp, "Hello, %s!", req.URL.Path[6:])
}

func formHandler(resp http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(resp, "ParseForm() err: %v", err)
	}

	fmt.Fprintf(resp, "Post request Successful")

	name := req.FormValue("name")
	address := req.FormValue("address")
	fmt.Fprintf(resp, "Name = %s\n", name)
	fmt.Fprintf(resp, "Address = %s\n", address)
}
