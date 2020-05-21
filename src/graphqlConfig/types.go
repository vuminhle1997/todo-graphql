package graphqlConfig

import (
	"github.com/graphql-go/graphql"
)

var TodoTypes = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Todo",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"title": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"done": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)
