package server

import (
	"os"

	"github.com/hesidoryn/swagger-example/router"
)

func Init() error {
	server := router.Load()
	port := os.Getenv("PORT")
	return server.Start(":" + port)
}
