package usecase

import (
	"encoding/json"
	"errors"
	"net/http"
	"project_bengkel/bengkel"
	"project_bengkel/entity"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func NewUsecaseModel(repos bengkel.RepoModelKendaraan) *repoModel {
	return &repoModel{
		repoMod: repos,
	}
}

type repoModel struct {
	repoMod bengkel.RepoModelKendaraan
}

func (rep repoModel) GetByNamaMerk(ctx *gin.Context) ([]entity.ModelKendaraan, error) {
	param := ctx.Param("model")

	if param == "" {
		return nil, errors.New(http.StatusText(http.StatusBadRequest))
	}

	result, err := rep.repoMod.GetByNamaMerk(param)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (rep repoModel) CreateModel(ctx *gin.Context) error {
	var model entity.ModelKendaraan
	year := time.Now().Year()

	err := ctx.ShouldBindJSON(&model)

	if err != nil {
		return err
	}

	tint, erro := strconv.ParseInt(model.Tahun, 10, 64)

	if erro != nil {
		return erro
	}

	if tint <= 1999 {
		return errors.New(http.StatusText(http.StatusBadRequest))
	}

	if tint > int64(year) {
		return errors.New(http.StatusText(http.StatusBadRequest))
	}

	err = rep.repoMod.CreateModel(model)

	if err != nil {
		return err
	}

	return nil
}

func (rep repoModel) UpdateModel(ctx *gin.Context) error {
	var model entity.ModelKendaraan
	year := time.Now().Year()
	id := ctx.Param("id")

	if id == "" {
		return errors.New(http.StatusText(http.StatusNotFound))
	}

	err := ctx.ShouldBindJSON(&model)

	if err != nil {
		return err
	}

	if model.IdMerk <= 0 {
		return errors.New(http.StatusText(http.StatusBadRequest))
	}

	tint, erro := strconv.ParseInt(model.Tahun, 10, 64)

	if tint <= 1999 {
		return errors.New(http.StatusText(http.StatusBadRequest))
	}

	if erro != nil {
		return erro
	}

	if tint > int64(year) {
		return errors.New(http.StatusText(http.StatusBadRequest))
	}

	err = rep.repoMod.UpdateModel(model, id)

	if err != nil {
		return err
	}
	return nil
}

func (rep repoModel) DeleteModel(ctx *gin.Context) error {
	var Inp struct {
		Id json.Number
	}

	if err := ctx.ShouldBindJSON(&Inp); err != nil {
		return err
	}

	id, err := Inp.Id.Int64()

	if err != nil {
		return errors.New(http.StatusText(http.StatusBadRequest))
	}

	err = rep.repoMod.DeleteModel(id)

	if err != nil {
		return err
	}

	return nil
}
