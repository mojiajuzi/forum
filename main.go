package main

import "net/http"

func main() {
	r := router()
	http.ListenAndServe(":8080", r)
}
