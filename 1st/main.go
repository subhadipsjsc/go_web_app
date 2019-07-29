package main

import (
	"fmt"
	"net/http"
)

func handleFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html> <h3> first Go server </h3> </html>")
}

func main() {
	fmt.Println("Hellow world")
	http.HandleFunc("/", handleFunction)
	http.ListenAndServe(":3000", nil)
}

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
// 	})

// 	http.ListenAndServe(":3000", nil)
// }
