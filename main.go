package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("We are alive !"))
	// fmt.Fprintf(w, "We're live!!!")
	fmt.Fprintf(w, "Hello %v, you are at %v, You have sent some query params too: %v. The method you used is %v", r.UserAgent(), r.URL.Path, r.URL.Query(), r.Method)

}
func main() {
	// http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/hello", HomePageHandler)
	fmt.Printf("Starting application on port %v\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
