package main

import (
	"net/http"
)

func main() {
	_, err := http.Get("http://localhost:3000/health")
	if err != nil {
		panic("could not access to running service")
	}
}
