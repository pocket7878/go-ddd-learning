package main

import (
	"log"

	"github.com/pocket7878/go-ddd-learning/infra"
	router "github.com/pocket7878/go-ddd-learning/presentation/controller/router"

	_ "github.com/mattn/go-sqlite3"
)

// @BasePath /
func main() {
	client, cleanupFunc, err := infra.BuildEntClient()
	defer cleanupFunc()

	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	r := router.BuildRouter(client)
	r.Run(":8080")
}
