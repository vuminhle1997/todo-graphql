package graphqlConfig

import (
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "RootQuery",
		Description: "Get data",
		Fields: graphql.Fields{
			"getToDoByID": &graphql.Field{
				Type:        TodoTypes,
				Description: "Get Todo item by its ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					fmt.Println(params.Info.RootValue.(map[string]interface{}))
					id, ok := params.Args["id"].(string)
					if ok {
						fmt.Println("ID in query", id)
						res := todoRepo.GetById(id)

						return res, nil
					}
					return nil, errors.New("404, Item not found")
				},
			},
			"getToDos": &graphql.Field{
				Type:        graphql.NewList(TodoTypes),
				Description: "Get a list of Todos",
				Args: graphql.FieldConfigArgument{
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					fmt.Println(params.Info)
					res := todoRepo.GetToDos()
					return res, nil
				},
			},
		},
	},
)
