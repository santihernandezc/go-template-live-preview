package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Listening...")
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/", fs)
	http.ListenAndServe(":8080", nil)
}
