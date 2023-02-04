package usecase

import (
	"encoding/json"
	"errors"
	"project_bengkel/bengkel"
	"project_bengkel/entity"

	"github.com/gin-gonic/gin"
)

func NewUsecaseSB(rep bengkel.RepoSB) *repoSb {
	return &repoSb{
		repo: rep,
	}
}

type repoSb struct {
	repo bengkel.RepoSB
}

func (rep repoSb) GetAllSB(ctx *gin.Context) ([]entity.ServiceBerkala, error) {
	result, err := rep.repo.GetAllSB()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (rep repoSb) CreateSB(ctx *gin.Context) error {
	var sb entity.ServiceBerkala

	ctx.ShouldBindJSON(&sb)

	if sb.Harga <= 10000 || sb.IdService <= 0 {
		return errors.New("bad request")
	}

	err := rep.repo.CreateSB(sb)

	if err != nil {
		return err
	}

	return nil
}

func (rep repoSb) UpdateSB(ctx *gin.Context) error {
	var sb entity.ServiceBerkala
	var param = ctx.Param("id")
	err := ctx.ShouldBindJSON(&sb)

	if param == "" {
		return errors.New("ID NOT FOUND")
	}

	if sb.Harga <= 10000 || sb.IdService <= 0 {
		return errors.New("bad request")
	}

	if err != nil {
		return err
	}

	err = rep.repo.UpdateSB(sb, param)

	if err != nil {
		return err
	}

	return nil
}

func (rep repoSb) DeleteSB(ctx *gin.Context) error {
	var inp struct {
		Id json.Number
	}

	err := ctx.ShouldBindJSON(&inp)

	if err != nil {
		return err
	}

	id, _ := inp.Id.Int64()

	err = rep.repo.DeleteSB(id)

	if err != nil {
		return err
	}

	return nil
}
