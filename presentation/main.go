package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pocket7878/go-ddd-learning/infra/ent"
	"github.com/pocket7878/go-ddd-learning/infra/ent/migrate"
	taskInfra "github.com/pocket7878/go-ddd-learning/infra/task"
	userInfra "github.com/pocket7878/go-ddd-learning/infra/user"
	"github.com/pocket7878/go-ddd-learning/presentation/docs"
	taskPresentation "github.com/pocket7878/go-ddd-learning/presentation/task"
	userPresentation "github.com/pocket7878/go-ddd-learning/presentation/user"
	taskUseCase "github.com/pocket7878/go-ddd-learning/usecase/task"
	userUseCase "github.com/pocket7878/go-ddd-learning/usecase/user"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/mattn/go-sqlite3"
)

// @BasePath /
func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	err = RunMigrations(*client)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	userController := BuildUserController(*client)
	taskController := BuildTaskController(*client)

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.POST("/users", userController.CreateUser)
	r.PATCH("/users/deactivated", userController.DeactivateUser)
	r.POST("/tasks", taskController.CreateTask)
	r.PATCH("/tasks/postponed", taskController.PosponeTask)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}

func RunMigrations(client ent.Client) error {
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

func BuildTaskController(client ent.Client) *taskPresentation.TaskController {
	taskRepo := taskInfra.NewTaskRdbRepository(client)
	taskCreateUseCase := taskUseCase.NewTaskCreateUseCase(taskRepo)
	taskPostponeUsecCase := taskUseCase.NewTaskPostponeUseCase(taskRepo)

	return taskPresentation.NewTaskController(*taskCreateUseCase, *taskPostponeUsecCase)
}

func BuildUserController(client ent.Client) *userPresentation.UserController {
	userRepo := userInfra.NewUserRdbRepository(client)
	userCreateUseCase := userUseCase.NewUserCreateUseCase(userRepo)
	userDeactivateUseCase := userUseCase.NewUserDeactivateUseCase(userRepo)

	return userPresentation.NewUserController(*userCreateUseCase, *userDeactivateUseCase)
}
