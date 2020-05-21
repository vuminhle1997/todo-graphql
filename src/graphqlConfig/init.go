package graphqlConfig

import (
	"log"
	"todo-graphql/src/models"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

var todos []models.Todo = []models.Todo{
	models.Todo{
		ID:          "42osu1",
		Title:       "Golang ToDo",
		Description: "Learn more Golang",
		Done:        false,
	},
	models.Todo{
		ID:          "0123ng",
		Title:       "NodeJS",
		Description: "Learn more Node",
		Done:        true,
	},
}

func Handler() gin.HandlerFunc {
	schema, err := InitSchema()
	if err != nil {
		log.Fatalf("failed")
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	log.Println("Handler succesfull created")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// TODO: Get Header later for JWT
// func ConfigHandler() gin.HandlerFunc {
// 	schema, err := InitSchema()
// 	if err != nil {
// 		log.Fatalf("failed")
// 	}

// 	h := handler.New(&handler.Config{
// 		Schema: &schema,
// 		Pretty: true,
// 		RootObjectFn: func(ctx *gin.Context, req *gin.Context.Request) map[string]interface{} {
// 			ctx := context.WithValue(c.Request.Context(), "header", c.Request.Header)
// 		},
// 	})

// 	return func(c *gin.Context) {

// 		h.ContextHandler(ctx, c.Writer, c.Request)
// 	}
// }
