package usecase

import (
	"errors"
	"project_bengkel/bengkel"
	"project_bengkel/entity"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

func NewUsecaseMerk(repos bengkel.RepoMerk) *repoMerk {
	rp := &repoMerk{
		repository: repos,
	}

	return rp
}

type repoMerk struct {
	repository bengkel.RepoMerk
}

func (rep repoMerk) GetAllMerk(ctx *gin.Context) ([]entity.MerkKendaraan, error) {
	result, err := rep.repository.GetAllMerk()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (rep repoMerk) CreateMerk(ctx *gin.Context) error {
	var data entity.MerkKendaraan
	err := ctx.ShouldBindJSON(&data)

	if err != nil {
		return err
	}

	err = rep.repository.CreateMerk(data)

	if err != nil {
		return err
	}

	return nil
}

func (rep repoMerk) UpdateMerk(ctx *gin.Context) error {
	var data entity.MerkKendaraan
	var id string = ctx.Param("id")
	var eror error

	if id == "" {
		eror = errors.New("ID Not Found")
		return eror
	}

	err := ctx.ShouldBindJSON(&data)

	if err != nil {
		return err
	}

	err = rep.repository.UpdateMerk(data, id)

	if err != nil {
		return err
	}

	return nil
}

func (rep repoMerk) DeleteMerk(ctx *gin.Context) error {
	var InputJson struct {
		Id json.Number
	}

	if err := ctx.ShouldBindJSON(&InputJson); err != nil {
		return err
	}

	id, _ := InputJson.Id.Int64()

	if id <= 0 {
		return errors.New("bad request")
	}

	err := rep.repository.DeleteMerk(id)

	if err != nil {
		return err
	}

	return nil
}
