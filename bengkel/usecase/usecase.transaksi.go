package usecase

import (
	"encoding/json"
	"errors"
	"project_bengkel/bengkel"
	"project_bengkel/entity/user"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func NewUsecaseTrans(rps bengkel.RepoTrans) reposT {
	return reposT{
		rep: rps,
	}
}

type reposT struct {
	rep bengkel.RepoTrans
}

// Transaksi Service
func (rp reposT) GetAllTserv(ctx *gin.Context) ([]user.TransaksiService, error) {
	result, err := rp.rep.GetAllTserv()

	if err != nil {
		return nil, err
	}

	return result, nil
}
func (rp reposT) CreateTserv(ctx *gin.Context) error {
	var tsv user.TransaksiService
	var tanggalLayout time.Time
	err := ctx.ShouldBindJSON(&tsv)

	if err != nil {
		return err
	}

	thn, _ := strconv.Atoi(tsv.TahunModel)

	if thn > time.Now().Year() {
		return errors.New("model invalid")
	}
	var cek []string = strings.Split(tsv.Tanggal, " ")
	if cek[2] == "AM" {
		lay := "2006-01-02 3:04 AM"
		tanggalLayout, _ = time.Parse(lay, tsv.Tanggal)
		tsv.TanggalService = tanggalLayout
	} else {
		lay := "2006-01-02 3:04 PM"
		tanggalLayout, _ = time.Parse(lay, tsv.Tanggal)
		tsv.TanggalService = tanggalLayout
	}
	if tsv.TanggalService.Before(time.Now()) {
		return errors.New("tanggal service invalid")
	}

	err = rp.rep.CreateTserv(tsv)

	if err != nil {
		return err
	}

	return nil
}

func (rp reposT) UpdateTserv(ctx *gin.Context) error {
	var tsv user.TransaksiService
	var id = ctx.Param("id")
	var tanggalLayout time.Time

	if id == "" {
		return errors.New("ID NOT FOUND")
	}

	err := ctx.ShouldBindJSON(&tsv)

	if err != nil {
		return err
	}

	thn, _ := strconv.Atoi(tsv.TahunModel)

	if thn > time.Now().Year() {
		return errors.New("model invalid")
	}

	var cek []string = strings.Split(tsv.Tanggal, " ")
	if cek[2] == "AM" {
		lay := "2006-01-02 3:04 AM"
		tanggalLayout, _ = time.Parse(lay, tsv.Tanggal)
		tsv.TanggalService = tanggalLayout
	} else {
		lay := "2006-01-02 3:04 PM"
		tanggalLayout, _ = time.Parse(lay, tsv.Tanggal)
		tsv.TanggalService = tanggalLayout
	}
	if tsv.TanggalService.Before(time.Now()) {
		return errors.New("tanggal service invalid")
	}

	err = rp.rep.UpdateTserv(tsv, id)

	if err != nil {
		return err
	}

	return nil
}

func (rp reposT) DeleteTserv(ctx *gin.Context) error {
	var inp struct {
		Id json.Number
	}

	if err := ctx.ShouldBindJSON(&inp); err != nil {
		return err
	}

	id, _ := inp.Id.Int64()

	if id <= 0 {
		return errors.New("ID NOT FOUND")
	}

	err := rp.rep.DeleteTserv(id)

	if err != nil {
		return err
	}

	return nil
}

// Transaksi SparePart
func (rp reposT) GetAllTsp(ctx *gin.Context) ([]user.TransaksiSparepart, error) {
	result, err := rp.rep.GetAllTsp()

	if err != nil {
		return nil, err
	}

	return result, nil
}
func (rp reposT) CreateTsp(ctx *gin.Context) error {
	var tsp user.TransaksiSparepart

	err := ctx.ShouldBindJSON(&tsp)

	if err != nil {
		return err
	}

	thn, _ := strconv.Atoi(tsp.TahunModel)

	if thn > time.Now().Year() {
		return errors.New("model invalid")
	}

	if tsp.Jumlah <= 0 {
		return errors.New("jumlah invalid")
	}

	err = rp.rep.CreateTsp(tsp)

	if err != nil {
		return err
	}

	return nil
}
func (rp reposT) UpdateTsp(ctx *gin.Context) error {
	var tsp user.TransaksiSparepart
	var id = ctx.Param("id")

	if id == "" {
		return errors.New("ID NOT FOUND")
	}

	err := ctx.ShouldBindJSON(&tsp)

	if err != nil {
		return err
	}

	thn, _ := strconv.Atoi(tsp.TahunModel)

	if thn > time.Now().Year() {
		return errors.New("model invalid")
	}

	if tsp.Jumlah <= 0 {
		return errors.New("jumlah invalid")
	}

	err = rp.rep.UpdateTsp(tsp, id)

	if err != nil {
		return err
	}

	return nil
}
func (rp reposT) DeleteTsp(ctx *gin.Context) error {
	var inp struct {
		Id json.Number
	}

	if err := ctx.ShouldBindJSON(&inp); err != nil {
		return err
	}

	id, _ := inp.Id.Int64()

	err := rp.rep.DeleteTsp(id)

	if err != nil {
		return err
	}

	return nil
}
