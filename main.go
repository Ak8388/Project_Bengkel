package main

import (
	"project_bengkel/connection"
	"project_bengkel/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Db := connection.Connection()

	rh := &router.Routes{
		Db: Db,
		R:  r,
	}
	rh.Routers()

	r.Run(":8080")
}
