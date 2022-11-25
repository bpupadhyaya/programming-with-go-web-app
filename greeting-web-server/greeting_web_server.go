package main

import (
	"fmt"
	"log"
	"net/http"
)

func greet(respWriter http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Println(req.Form)
	fmt.Println("path: ", req.URL.Path)
	fmt.Println("scheme: ", req.URL.Scheme)
	fmt.Println("URL long: ", req.Form["url_long"])
	fmt.Println("--------")
	for key, value := range req.Form {
		fmt.Println("key: ", key)
		fmt.Println("value: ", value)
	}
	fmt.Fprintf(respWriter, "Message from the server")
}

func main() {
	http.HandleFunc("/", greet)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Error from ListenAndServe: ", err)
	}
}
