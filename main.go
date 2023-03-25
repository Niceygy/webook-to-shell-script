package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func getScript() string {
	argsInt := os.Args[1:]
	if argsInt == nil {
		fmt.Println("No arguments passed! Please inclue location of script")
		os.Exit(1)
	}
	argsOut := strings.Join(argsInt, " ")
	return argsOut
}

func webooktoshellscript() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "All working!, %q", r.URL.Path)
	})
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", r.URL.Path)
		cmd := exec.Command(getScript())
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Running!")
		}
	})
	log.Fatal(http.ListenAndServe(":7002", nil))
}

func main() {
	webooktoshellscript()
}
