package main

import (
	"context"
	"log"

	"github.com/pocket7878/go-ddd-learning/infra/ent"
	"github.com/pocket7878/go-ddd-learning/infra/ent/migrate"
	router "github.com/pocket7878/go-ddd-learning/presentation/controller/router"

	_ "github.com/mattn/go-sqlite3"
)

// @BasePath /
func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	err = RunMigrations(client)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	r := router.BuildRouter(client)
	r.Run(":8080")
}

func RunMigrations(client *ent.Client) error {
	ctx := context.Background()
	// マイグレーションの実行
	err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		return err
	}

	return nil
}
