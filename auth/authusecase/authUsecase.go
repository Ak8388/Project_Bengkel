package authusecase

import (
	"project_bengkel/auth"
	"project_bengkel/entity/user"
	"project_bengkel/jwt"

	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)

type reposAuthUsecase struct {
	repo auth.AuthRepos
}

func NewUsecaseAuth(rep auth.AuthRepos) reposAuthUsecase {
	return reposAuthUsecase{
		repo: rep,
	}
}

func (us reposAuthUsecase) Register(ctx *gin.Context) error {
	var data user.Regist
	if err := ctx.ShouldBindJSON(&data); err != nil {
		return err
	}

	hassPass, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)

	if err != nil {
		return err
	}

	data.Password = string(hassPass)

	if err = us.repo.Register(data); err != nil {
		return err
	}

	return nil
}
func (us reposAuthUsecase) Login(ctx *gin.Context) error {
	var data auth.Login

	if err := ctx.ShouldBindJSON(&data); err != nil {
		return err
	}

	result, err := us.repo.Login(data.Username)

	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(data.Password))

	if err != nil {
		return err
	}

	token, errTok := jwt.Token(int(result.ID))

	if errTok != nil {
		return errTok
	}
	ctx.SetCookie("token", token.AcToken, 3600, "/", "localhost", false, true)

	return nil
}
func (us reposAuthUsecase) Logout(ctx *gin.Context) error {
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
	return nil
}
