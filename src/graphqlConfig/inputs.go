package graphqlConfig

import (
	"github.com/graphql-go/graphql"
)

var todoInputType *graphql.InputObject = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "TodoInputs",
		Fields: graphql.InputObjectConfigFieldMap{
			"title": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"description": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"done": &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
		},
	},
)

var todoInputArgs = graphql.FieldConfigArgument{
	"todoInput": &graphql.ArgumentConfig{
		Type: todoInputType,
	},
}

var todoInputArgsWithID = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"todoInput": &graphql.ArgumentConfig{
		Type: todoInputType,
	},
}
