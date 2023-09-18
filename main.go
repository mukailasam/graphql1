package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	app := NewGraphql()

	http.HandleFunc("/graphql", app.EntryPoint)

	fmt.Println("Listen on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
}
