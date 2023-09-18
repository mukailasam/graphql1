package resolvers

import (
	"database/sql"
	"errors"

	//"fmt"

	"github.com/ftsog/graphql1/models"
	"github.com/graphql-go/graphql"
)

var invalidQuery = errors.New("Invalid Query")
var internalServerError = errors.New("Internal Server Error")
var notFound = errors.New("Resources NotFound")

type messagae struct {
	msg string
}

type ResolverInterface interface {
	CreateMessage(title, body string) error
	ReadMessage(id int) (interface{}, error)
	ReadMessages() (interface{}, error)
	UpdateMessage(id int, title, body string) error
	DeleteMessage(id int) error
	PostExist(id int) bool
}
type Resolver struct {
	Model ResolverInterface
}

func (r Resolver) PostQueryResolver(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	if ok {
		p, err := r.Model.ReadMessage(id)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, notFound
			}

			return nil, internalServerError
		}
		np := p.(models.Post)
		return np, nil
	}

	return nil, invalidQuery
}

func (r Resolver) PostsQueryResolver(p graphql.ResolveParams) (interface{}, error) {
	mp, err := r.Model.ReadMessages()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, notFound
		}
		return nil, internalServerError

	}

	nmp := mp.([]models.Post)
	return nmp, nil
}

func (r Resolver) CreatePostMutateResolver(p graphql.ResolveParams) (interface{}, error) {
	title, ok1 := p.Args["title"].(string)
	body, ok2 := p.Args["body"].(string)
	if ok1 != true && ok2 != true {
		return nil, invalidQuery
	}

	err := r.Model.CreateMessage(title, body)
	if err != nil {
		return nil, internalServerError
	}

	return nil, nil

}

func (r Resolver) UpdatePostMutateResolver(p graphql.ResolveParams) (interface{}, error) {
	id, ok0 := p.Args["id"].(int)
	title, ok1 := p.Args["title"].(string)
	body, ok2 := p.Args["body"].(string)
	if ok0 != true && ok1 != true && ok2 != true {
		return nil, invalidQuery
	}

	Exist := r.Model.PostExist(id)
	if !Exist {
		return nil, notFound
	}

	err := r.Model.UpdateMessage(id, title, body)
	if err != nil {
		return nil, internalServerError
	}

	return nil, nil
}

func (r Resolver) DeletePostMutateResolver(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	if !ok {
		return nil, invalidQuery
	}

	Exist := r.Model.PostExist(id)
	if !Exist {
		return nil, notFound
	}

	err := r.Model.DeleteMessage(id)
	if err != nil {
		return nil, internalServerError
	}

	return nil, nil
}
