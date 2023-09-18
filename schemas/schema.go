package schemas

import (
	"github.com/graphql-go/graphql"
)

type SchemaInterface interface {
	PostQueryResolver(p graphql.ResolveParams) (interface{}, error)
	PostsQueryResolver(p graphql.ResolveParams) (interface{}, error)
	CreatePostMutateResolver(p graphql.ResolveParams) (interface{}, error)
	UpdatePostMutateResolver(p graphql.ResolveParams) (interface{}, error)
	DeletePostMutateResolver(p graphql.ResolveParams) (interface{}, error)
}

type Schema struct {
	Resolver SchemaInterface
}

func (s Schema) PostType() *graphql.Object {
	postType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id":    &graphql.Field{Type: graphql.Int},
			"title": &graphql.Field{Type: graphql.String},
			"body":  &graphql.Field{Type: graphql.String},
		},
	})

	return postType
}

func (s Schema) PostQueryType() *graphql.Object {

	postType := s.PostType()
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"post": &graphql.Field{
				Type:        postType,
				Description: "Get a single post",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},

				Resolve: s.Resolver.PostQueryResolver,
			},

			"posts": &graphql.Field{
				Type:        graphql.NewList(postType),
				Description: "Get all posts",
				Resolve:     s.Resolver.PostsQueryResolver,
			},
		},
	})

	return queryType
}

func (s Schema) MutationType() *graphql.Object {
	postType := s.PostType()
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"create": &graphql.Field{
				Type:        postType,
				Description: "create post",
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},

					"body": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: s.Resolver.CreatePostMutateResolver,
			},

			"update": &graphql.Field{
				Type:        postType,
				Description: "update post",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},

					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},

					"body": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},

				Resolve: s.Resolver.UpdatePostMutateResolver,
			},

			"delete": &graphql.Field{
				Type:        postType,
				Description: "delete post",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},

				Resolve: s.Resolver.DeletePostMutateResolver,
			},
		},
	})

	return mutationType

}
