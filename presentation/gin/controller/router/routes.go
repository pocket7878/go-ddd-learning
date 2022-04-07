package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pocket7878/go-ddd-learning/infra/ent"
	taskInfra "github.com/pocket7878/go-ddd-learning/infra/task"
	userInfra "github.com/pocket7878/go-ddd-learning/infra/user"
	taskPresentation "github.com/pocket7878/go-ddd-learning/presentation/gin/controller/task"
	userPresentation "github.com/pocket7878/go-ddd-learning/presentation/gin/controller/user"
	"github.com/pocket7878/go-ddd-learning/presentation/gin/docs"
	taskUseCase "github.com/pocket7878/go-ddd-learning/usecase/task"
	userUseCase "github.com/pocket7878/go-ddd-learning/usecase/user"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/mattn/go-sqlite3"
)

func BuildRouter(client *ent.Client) *gin.Engine {
	userController := buildUserController(client)
	taskController := buildTaskController(client)

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.POST("/users", userController.CreateUser)
	r.PATCH("/users/deactivated", userController.DeactivateUser)
	r.POST("/tasks", taskController.CreateTask)
	r.PATCH("/tasks/postponed", taskController.PosponeTask)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}

func buildTaskController(client *ent.Client) *taskPresentation.TaskController {
	taskRepo := taskInfra.NewTaskRdbRepository(client)
	taskCreateUseCase := taskUseCase.NewTaskCreateUseCase(taskRepo)
	taskPostponeUsecCase := taskUseCase.NewTaskPostponeUseCase(taskRepo)

	return taskPresentation.NewTaskController(*taskCreateUseCase, *taskPostponeUsecCase)
}

func buildUserController(client *ent.Client) *userPresentation.UserController {
	userRepo := userInfra.NewUserRdbRepository(client)
	userCreateUseCase := userUseCase.NewUserCreateUseCase(userRepo)
	userDeactivateUseCase := userUseCase.NewUserDeactivateUseCase(userRepo)

	return userPresentation.NewUserController(*userCreateUseCase, *userDeactivateUseCase)
}
