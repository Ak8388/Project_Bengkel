package usecase

import (
	"encoding/json"
	"errors"
	"project_bengkel/bengkel"
	"project_bengkel/entity"

	"github.com/gin-gonic/gin"
)

func NewUsecaseKat(rep bengkel.RepoKat) *repoKat {
	return &repoKat{
		repo: rep,
	}
}

type repoKat struct {
	repo bengkel.RepoKat
}

func (rep repoKat) GetAllKat(ctx *gin.Context) ([]entity.Kategori, error) {
	result, err := rep.repo.GetAllKat()

	if err != nil {
		return nil, err
	}

	return result, nil
}
func (rep repoKat) CreateKat(ctx *gin.Context) error {
	var kategori entity.Kategori
	err := ctx.ShouldBindJSON(&kategori)

	if err != nil {
		return err
	}

	err = rep.repo.CreateKat(kategori)

	if err != nil {
		return err
	}

	return nil
}

func (rep repoKat) UpdateKat(ctx *gin.Context) error {
	var kategori entity.Kategori
	var id = ctx.Param("id")

	if id == "" {
		return errors.New("ID NOT FOUND")
	}

	err := rep.repo.UpdateKat(kategori, id)

	if err != nil {
		return err
	}

	return nil
}

func (rep repoKat) DeleteKat(ctx *gin.Context) error {
	var inputkat struct {
		Id json.Number
	}

	err := ctx.ShouldBindJSON(&inputkat)

	if err != nil {
		return err
	}

	id, _ := inputkat.Id.Int64()

	err = rep.repo.DeleteKat(id)

	if err != nil {
		return err
	}

	return nil

}
