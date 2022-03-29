package user

import (
	"github.com/gin-gonic/gin"
	userDomain "github.com/pocket7878/go-ddd-learning/domain/user"
)

type DeactivateUserParams struct {
	ID string `json:"id"`
}

// @Schemes
// @Description create User
// @Accept json
// @Produce json
// @Tags user
// @Param deactivateUserParam body DeactivateUserParams true "deactivate user"
// @Success 200 {object} user.UserJsonResponseBody
// @Router /users/deactivated [patch]
func (u *UserController) DeactivateUser(ctx *gin.Context) {
	var p DeactivateUserParams
	err := ctx.BindJSON(&p)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userID := userDomain.ReconstructUserID(p.ID)
	user, err := u.userDeactivateUseCase.DeactivateUser(ctx, *userID)
	if err != nil {
		ctx.JSON(422, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, NewUserJsonResponseBody(user))
}
