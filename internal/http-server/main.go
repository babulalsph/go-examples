package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Simple http server example server request and response data
func main() {

	log.Println("Server starting...")
	http.HandleFunc("/hello", handlerFunc)

	log.Println("Please connect to localhost:8080/hello")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func handlerFunc(writer http.ResponseWriter, req *http.Request) {

	// log `Hello word` message on server
	log.Println("Hello word")

	// read request url
	log.Println("Request URL: ", req.URL)

	// read request body data
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(writer, "Oops invalid request", http.StatusBadRequest)
		return
	}

	// write response back to write handler
	fmt.Fprintf(writer, "Response data : %s", data)
}
