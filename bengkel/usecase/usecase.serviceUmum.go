package usecase

import (
	"errors"
	"project_bengkel/bengkel"
	"project_bengkel/entity"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

func NewUsecaseSU(repoSu bengkel.RepoSU) *repos {
	return &repos{
		rep: repoSu,
	}
}

type repos struct {
	rep bengkel.RepoSU
}

func (rp repos) GetAllSU(ctx *gin.Context) ([]entity.ServiceUmum, error) {
	result, err := rp.rep.GetAllSU()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (rp repos) CreateSU(ctx *gin.Context) error {
	var su entity.ServiceUmum

	ctx.ShouldBindJSON(&su)

	if su.Harga <= 10000 || su.IdService <= 0 {
		return errors.New("bad request")
	}

	err := rp.rep.CreateSU(su)

	if err != nil {
		return err
	}

	return nil
}

func (rp repos) UpdateSU(ctx *gin.Context) error {
	var su entity.ServiceUmum
	var param = ctx.Param("id")
	err := ctx.ShouldBindJSON(&su)

	if param == "" {
		return errors.New("ID NOT FOUND")
	}

	if su.Harga <= 10000 || su.IdService <= 0 {
		return errors.New("bad request")
	}

	if err != nil {
		return err
	}

	err = rp.rep.UpdateSU(su, param)

	if err != nil {
		return err
	}

	return nil
}

func (rp repos) DeleteSU(ctx *gin.Context) error {
	var inp struct {
		Id json.Number
	}

	err := ctx.ShouldBindJSON(&inp)

	if err != nil {
		return err
	}

	id, _ := inp.Id.Int64()

	err = rp.rep.DeleteSU(id)

	if err != nil {
		return err
	}

	return nil

}
