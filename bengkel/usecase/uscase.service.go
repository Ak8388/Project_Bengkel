package usecase

import (
	"errors"
	"project_bengkel/bengkel"
	"project_bengkel/entity"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

func NewUsecaseService(repos bengkel.RepoService) repo {
	rp := repo{
		repository: repos,
	}

	return rp
}

type repo struct {
	repository bengkel.RepoService
}

func (rep repo) GetAllService(ctx *gin.Context) ([]entity.Service, error) {
	result, err := rep.repository.GetAllService()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (rep repo) CreateService(ctx *gin.Context) error {
	var data entity.Service
	err := ctx.ShouldBindJSON(&data)

	if err != nil {
		return err
	}

	err = rep.repository.CreateService(data)

	if err != nil {
		return err
	}

	return nil
}

func (rep repo) UpdateService(ctx *gin.Context) error {
	var data entity.Service
	var id string = ctx.Param("id")

	if id == "" {
		return errors.New("ID Not Found")
	}

	err := ctx.ShouldBindJSON(&data)

	if err != nil {
		return err
	}

	err = rep.repository.UpdateService(data, id)

	if err != nil {
		return err
	}

	return nil
}

func (rep repo) DeleteService(ctx *gin.Context) error {
	var InputJson struct {
		Id json.Number
	}

	if err := ctx.ShouldBindJSON(&InputJson); err != nil {
		return err
	}

	id, err := InputJson.Id.Int64()

	if err != nil {
		return err
	}

	if id <= 0 {
		return errors.New("bad request")
	}

	err = rep.repository.DeleteService(id)

	if err != nil {
		return err
	}

	return nil
}
