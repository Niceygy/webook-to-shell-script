package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func webooktoshellscript() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", r.URL.Path)
	})
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", r.URL.Path)
		cmd := exec.Command("/bin/bash", "script.sh")
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	webooktoshellscript()
}
