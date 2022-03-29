package user

import (
	"github.com/gin-gonic/gin"
	userDomain "github.com/pocket7878/go-ddd-learning/domain/user"
	userUseCase "github.com/pocket7878/go-ddd-learning/usecase/user"
)

type UserController struct {
	userCreateUseCase     userUseCase.UserCreateUseCase
	userDeactivateUseCase userUseCase.UserDeactivateUseCase
}

func NewUserController(userCreateUseCase userUseCase.UserCreateUseCase, userDeactivateUseCase userUseCase.UserDeactivateUseCase) *UserController {
	return &UserController{
		userCreateUseCase:     userCreateUseCase,
		userDeactivateUseCase: userDeactivateUseCase,
	}
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	userName := ctx.PostForm("name")
	user, err := u.userCreateUseCase.CreateUser(ctx, userName)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, NewUserJsonResponseBody(user))
}

func (u *UserController) DeactivateUser(ctx *gin.Context) {
	userIDString := ctx.PostForm("id")
	userID := userDomain.ReconstructUserID(userIDString)
	user, err := u.userDeactivateUseCase.DeactivateUser(ctx, *userID)
	if err != nil {
		ctx.JSON(422, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, NewUserJsonResponseBody(user))
}
