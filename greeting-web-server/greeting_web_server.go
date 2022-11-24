package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func greet(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.ParseForm()
	fmt.Println(respWriter.Form)
	fmt.Println("path: ", respWriter.URL.Path)
	fmt.Println("scheme: ", respWriter.URL.Scheme)
	fmt.Println("URL long: ", respWriter.Form["url_long"])
	for key, value := range respWriter.Form {
		fmt.Printf("key: %s, value: %s", key, value)
	}
}

func main() {
	http.HandleFunc("/", greet)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error from ListenAndServe: ", err)
	}
}
