package auth

import (
	"project_bengkel/entity/user"

	"github.com/gin-gonic/gin"
)

type AuthUsecase interface {
	Register(ctx *gin.Context) error
	Login(ctx *gin.Context) error
	Logout(ctx *gin.Context) error
}

type AuthRepos interface {
	Register(user user.Regist) error
	Login(string) (*user.Regist, error)
	Logout() error
}
