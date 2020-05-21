package graphqlConfig

import (
	"fmt"
	"todo-graphql/src/models"
	"todo-graphql/src/repository"

	"github.com/graphql-go/graphql"
)

var todoRepo = repository.New()

var RootMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "RootMutation",
		Description: "Muation type for ToDo",
		Fields: graphql.Fields{
			"createToDo": &graphql.Field{
				Type:        TodoTypes,
				Description: "Create ToDo item",
				Args:        todoInputArgs,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					fmt.Println(params.Context)
					var inputs map[string]interface{} = params.Args["todoInput"].(map[string]interface{})

					// fmt.Println(inputs, reflect.TypeOf(inputs))

					var title string = inputs["title"].(string)
					var description string = inputs["description"].(string)
					var done bool = inputs["done"].(bool)

					todo := models.Todo{
						Title:       title,
						Description: description,
						Done:        done,
					}

					fmt.Println(todo)

					res := todoRepo.Save(todo)

					return res, nil
				},
			},
			"updateToDo": &graphql.Field{
				Type:        TodoTypes,
				Description: "Update ToDo item by its ID",
				Args:        todoInputArgsWithID,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					fmt.Println(params.Info)
					var id string = params.Args["id"].(string)
					var inputs map[string]interface{} = params.Args["todoInput"].(map[string]interface{})

					todoInput := repository.TodoInput{
						Title:       inputs["title"].(string),
						Description: inputs["description"].(string),
						Done:        inputs["done"].(bool),
					}

					res := todoRepo.UpdateById(id, todoInput)

					return res, nil
				},
			},
		},
	},
)
