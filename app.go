package main

import (
	"fmt"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Ahoj svete z Go! 🦄🦄🦄 %s\n", hostname)
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server started.")
	http.ListenAndServe(":80", nil)
}
