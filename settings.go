package main

import (
	"github.com/ftsog/graphql1/db"
	"github.com/ftsog/graphql1/models"
	"github.com/ftsog/graphql1/resolvers"
	"github.com/ftsog/graphql1/schemas"
)

func NewGraphql() Application {
	dbConn := db.DatabaseConnection()

	model := models.Model{
		DB: dbConn,
	}

	resolver := resolvers.Resolver{
		Model: model,
	}

	schema := schemas.Schema{
		Resolver: resolver,
	}

	app := Application{
		Schema: schema,
	}

	return app

}
