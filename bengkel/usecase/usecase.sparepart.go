package usecase

import (
	"encoding/json"
	"errors"
	"project_bengkel/bengkel"
	"project_bengkel/entity"

	"github.com/gin-gonic/gin"
)

func NewUsecaseSP(repo bengkel.RepoSparePart) *repoSp {
	return &repoSp{
		rep: repo,
	}
}

type repoSp struct {
	rep bengkel.RepoSparePart
}

func (repo repoSp) GetAllSP(ctx *gin.Context) ([]entity.SparePart, error) {
	result, err := repo.rep.GetAllSP()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo repoSp) CreateSP(ctx *gin.Context) error {
	var sp entity.SparePart

	err := ctx.ShouldBindJSON(&sp)

	if err != nil {
		return err
	}

	if sp.IdKat <= 0 || sp.IdModel <= 0 {
		return errors.New("ID Kategoris or ID Model Not Valid")
	}

	if sp.Harga <= 1000 || sp.Stok <= 0 {
		if sp.Harga <= 1000 {
			return errors.New("harga invalid")
		} else {
			return errors.New("stok invalid")
		}
	}

	if sp.Nama == " " {
		return errors.New("name invalid")
	}

	err = repo.rep.CreateSP(sp)

	if err != nil {
		return err
	}
	return nil

}

func (repo repoSp) UpdateSP(ctx *gin.Context) error {
	var sp entity.SparePart
	var id = ctx.Param("id")

	if id == "" {
		return errors.New("please input sparepart ID")
	}

	err := ctx.ShouldBindJSON(&sp)

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	if sp.IdKat <= 0 || sp.IdModel <= 0 {
		return errors.New("ID kategoris or ID model invalid")
	}

	if sp.Harga <= 1000 || sp.Stok <= 0 {
		if sp.Harga <= 1000 {
			return errors.New("harga invalid")
		} else {
			return errors.New("stok invalid")
		}
	}

	if sp.Nama == " " {
		return errors.New("name invalid")
	}

	err = repo.rep.UpdateSP(sp, id)

	if err != nil {
		return err
	}

	return nil
}

func (repo repoSp) DeleteSP(ctx *gin.Context) error {
	var inp struct {
		Id json.Number
	}

	err := ctx.ShouldBindJSON(&inp)

	if err != nil {
		return err
	}

	id, _ := inp.Id.Int64()

	if id <= 0 {
		return errors.New("please input ID sparepart")
	}

	err = repo.rep.DeleteSP(id)

	if err != nil {
		return err
	}

	return nil
}
