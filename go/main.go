package main

import (
	"fmt"
	"net/http"
)

func greeting(text string) string {
	return fmt.Sprintf("<b>%s</b>", text)
}

func handleGreeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, greeting("Code.education Rocks!"))
}

func main() {
	fmt.Printf("Starting server at port 8000\n")
	http.HandleFunc("/", handleGreeting)
	http.ListenAndServe(":8000", nil)
}