package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

type HandlerInterface interface {
	RootSchema() graphql.Schema
}

type Application struct {
	Schema HandlerInterface
}

func (app Application) EntryPoint(w http.ResponseWriter, r *http.Request) {
	schema := app.Schema.RootSchema()
	requestString := r.URL.Query().Get("query")
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: requestString,
	})

	if result.HasErrors() {
		fmt.Println(result.Errors)
	}

	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		fmt.Println(err)
	}

}
