package infra

import (
	"context"

	"github.com/pocket7878/go-ddd-learning/infra/ent"
	"github.com/pocket7878/go-ddd-learning/infra/ent/migrate"

	_ "github.com/mattn/go-sqlite3"
)

func BuildEntClient() (*ent.Client, func(), error) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return nil, nil, err
	}

	err = runMigrations(client)
	if err != nil {
		return nil, nil, err
	}

	return client, func() { client.Close() }, nil
}

func runMigrations(client *ent.Client) error {
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
