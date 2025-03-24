package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Finally the setup is done.")
}

func main() {
	http.HandleFunc("/", handler)
	port := ":8080"
	fmt.Println("Server is running on port", port)
	http.ListenAndServe(port, nil)
}
