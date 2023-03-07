package authhandler

import (
	"net/http"
	"project_bengkel/auth"

	"github.com/gin-gonic/gin"
)

func NewHandlerAuth(us auth.AuthUsecase, r *gin.RouterGroup) {
	hand := UsecaseAuthHandler{
		uscase: us,
	}

	v2 := r.Group("Auth")
	v2.POST("regist", hand.Register)
	v2.POST("login", hand.Login)
	v2.GET("logout", hand.Logout)
}

type UsecaseAuthHandler struct {
	uscase auth.AuthUsecase
}

func (us UsecaseAuthHandler) Register(ctx *gin.Context) {
	err := us.uscase.Register(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Regist"})
}
func (us UsecaseAuthHandler) Login(ctx *gin.Context) {
	err := us.uscase.Login(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Login"})
}
func (us UsecaseAuthHandler) Logout(ctx *gin.Context) {
	err := us.uscase.Logout(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Logout Succses"})
}
