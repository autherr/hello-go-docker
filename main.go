package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Go Docker"))

	fmt.Fprintf(w, "\nEnvironment:\n")
	for _, env := range os.Environ() {
		fmt.Fprintf(w, "%s\n", env)
	}
	fmt.Fprintf(w, "\n")

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}
	fmt.Fprintf(w, "Generated on %s\n", hostname)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func printUser() {
	uid := strconv.Itoa(os.Getuid())

	passwd, err := ioutil.ReadFile("/etc/passwd")
	if err != nil {
		fmt.Printf("Error reading /etc/passwd: %v", err)
		return
	}
	for _, line := range strings.Split(string(passwd), "\n") {
		if len(line) == 0 {
			continue
		}
		entries := strings.Split(line, ":")
		name, id := entries[0], entries[2]
		if id == uid {
			fmt.Println("User:", name)
			fmt.Println("Uid:", id)
			return
		}
	}
}

func main() {
	printUser()
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/health", health)
	fmt.Println("listening on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
