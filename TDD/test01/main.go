package main

import (
	"fmt"
	"net/http"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "jackfrow")
}

func main() {
	fmt.Println("ListenAndServe :5000")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
