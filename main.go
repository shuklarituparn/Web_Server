package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return

	} // /hello is the get method

	fmt.Fprintf(w, "hello!")

} //w is the repsonse the server send, r is the request the user sends

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform error: %v", err)
		return
	}
	fmt.Fprintf(w, "Post request successfully")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name= %s\n", name)
	fmt.Fprintf(w, "Address= %s\n", address)

}
func main() {
	fileserver := http.FileServer(http.Dir("./static")) //php, go, nodejs look at the index file
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("Starting the server at the port 8080\n")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
