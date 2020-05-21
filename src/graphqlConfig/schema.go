package graphqlConfig

import (
	"log"

	"github.com/graphql-go/graphql"
)

func InitSchema() (graphql.Schema, error) {
	schameConfig := graphql.SchemaConfig{
		Query:    RootQuery,
		Mutation: RootMutation,
	}
	schema, err := graphql.NewSchema(schameConfig)

	if err != nil {
		log.Fatalf("failed to init schema")
	}

	log.Println("Schema successful created")
	return schema, nil
}
