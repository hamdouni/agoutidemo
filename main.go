package main

import (
	"log"
	"net/http"
	"strconv"
)

func StartApp(p int) {
	a := ":" + strconv.Itoa(p)
	log.Fatal(http.ListenAndServe(a, http.FileServer(http.Dir("./"))))
}
func main() {
	// Simple static webserver:
	StartApp(8888)
}
