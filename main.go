package main

import (
	"todo-graphql/src/graphqlConfig"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/graphql", graphqlConfig.Handler())
	r.Run()
}
