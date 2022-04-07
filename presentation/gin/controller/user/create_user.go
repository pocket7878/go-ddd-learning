package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type CreateUserParams struct {
	Name string `json:"name"`
}

// @Schemes
// @Description create User
// @Accept json
// @Produce json
// @Tags user
// @Param createUserParam body CreateUserParams true "create user"
// @Success 200 {object} user.UserJsonResponseBody
// @Router /users [post]
func (u *UserController) CreateUser(ctx *gin.Context) {
	var p CreateUserParams
	err := ctx.BindJSON(&p)
	if err != nil {
		ctx.JSON(400, gin.H{"error": fmt.Sprintf("Failed to read parameter: %s", err.Error())})
		return
	}

	user, err := u.userCreateUseCase.CreateUser(ctx, p.Name)
	if err != nil {
		ctx.JSON(422, gin.H{"error": fmt.Sprintf("Failed to create user: %s", err.Error())})
		return
	}

	ctx.JSON(201, NewUserJsonResponseBody(user))
}
