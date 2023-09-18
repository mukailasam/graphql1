package schemas

import (
	//"fmt"

	"github.com/graphql-go/graphql"
)

func (s Schema) RootSchema() graphql.Schema {
	queryType := s.PostQueryType()
	mutationType := s.MutationType()
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})

	return schema
}
