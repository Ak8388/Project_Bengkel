package repository

import (
	"errors"
	"fmt"
	"project_bengkel/entity"
	"project_bengkel/entity/user"

	"gorm.io/gorm"
)

func NewRepoTrans(db *gorm.DB) dbRepos {
	return dbRepos{
		Db: db,
	}
}

type dbRepos struct {
	Db *gorm.DB
}

// Transaksi Service
func (db dbRepos) GetAllTserv() ([]user.TransaksiService, error) {
	var transServ []user.TransaksiService
	err := db.Db.Find(&transServ).Error

	if err != nil {
		return nil, err
	}

	return transServ, nil
}

func (db dbRepos) CreateTserv(tsp user.TransaksiService) error {
	var serviceBk entity.ServiceBerkala
	var serviceUm entity.ServiceUmum
	var model []entity.ModelKendaraan
	var cek = 0

	if err := db.Db.First(&serviceBk, "service = ?", tsp.Service).Error; err != nil {
		if err = db.Db.First(&serviceUm, "service = ?", tsp.Service).Error; err != nil {
			return errors.New("service invalid")
		} else {
			tsp.TotalHarga = serviceUm.Harga
		}
	}

	tsp.TotalHarga = serviceBk.Harga

	db.Db.Find(&model, "model = ?", tsp.Model)

	for a := 0; a < len(model); a++ {
		if model[a].Model == tsp.Model {
			if model[a].Tahun == tsp.TahunModel {
				fmt.Println("Hallo")
				break
			}
		}
		cek++
	}

	if cek == len(model) {
		return errors.New("model not found")
	}

	db.Db.Create(&tsp)
	return nil
}

func (db dbRepos) UpdateTserv(tsp user.TransaksiService, param string) error {
	var serviceBk entity.ServiceBerkala
	var serviceUm entity.ServiceUmum
	var cek = 0
	var model []entity.ModelKendaraan

	if err := db.Db.First(&serviceBk, "service = ?", tsp.Service).Error; err != nil {
		if err = db.Db.First(&serviceUm, "service = ?", tsp.Service).Error; err != nil {
			return errors.New("service invalid")
		} else {
			tsp.TotalHarga = serviceUm.Harga
		}
	}

	tsp.TotalHarga = serviceBk.Harga

	db.Db.Find(&model, "model = ?", tsp.Model)

	for a := 0; a < len(model); a++ {
		if model[a].Model == tsp.Model {
			if model[a].Tahun == tsp.TahunModel {
				break
			}
		}
		cek++
	}

	if cek == len(model) {
		return errors.New("model not found")
	}

	err := db.Db.Where("id = ?", param).Updates(&tsp).Error

	if err != nil {
		return err
	}

	return nil
}

func (db dbRepos) DeleteTserv(id int64) error {

	if err := db.Db.Where("id = ?", id).Delete(&user.TransaksiService{}).Error; err != nil {
		return err
	}

	return nil
}

// Transaksi Sparepart
func (db dbRepos) GetAllTsp() ([]user.TransaksiSparepart, error) {
	var tsp []user.TransaksiSparepart
	err := db.Db.Find(&tsp).Error

	if err != nil {
		return nil, err
	}

	return tsp, nil
}

func (db dbRepos) CreateTsp(tsp user.TransaksiSparepart) error {
	var sparepart []entity.SparePart
	var model []entity.ModelKendaraan
	var id uint
	var cek = 0
	var x = 0

	db.Db.Find(&model, "model = ?", tsp.Model)

	for a := 0; a < len(model); a++ {
		if model[a].Model == tsp.Model {
			if model[a].Tahun == tsp.TahunModel {
				id = model[a].ID
				break
			}
		}
		cek++
	}

	if cek == len(model) {
		return errors.New("model not found")
	}

	cek = 0

	db.Db.Find(&sparepart, "id_model = ?", id)

	for a := 0; a < len(sparepart); a++ {
		if sparepart[a].IdModel == id {
			if sparepart[a].Nama == tsp.SparePart {
				x = a
				break
			}
		}
		cek++
	}

	if cek == len(sparepart) {
		return errors.New("spare part not found")
	}

	if tsp.Jumlah > sparepart[x].Stok {
		return errors.New("spare parts stock is running out")
	}
	tsp.TotalHarga = sparepart[x].Harga * tsp.Jumlah
	sparepart[x].Stok -= tsp.Jumlah
	fmt.Println(sparepart[x])
	fmt.Println(sparepart[x].ID)
	db.Db.Where("id = ?", sparepart[x].ID).Updates(&sparepart[x])
	db.Db.Create(&tsp)

	return nil
}

func (db dbRepos) UpdateTsp(tsp user.TransaksiSparepart, param string) error {
	var sparepart []entity.SparePart
	var model []entity.ModelKendaraan
	var id uint
	var cek = 0
	var x = 0
	var val user.TransaksiSparepart

	db.Db.Find(&model, "model = ?", tsp.Model)

	for a := 0; a < len(model); a++ {
		if model[a].Model == tsp.Model {
			if model[a].Tahun == tsp.TahunModel {
				id = model[a].ID
				break
			}
		}
		cek++
	}

	if cek == len(model) {
		return errors.New("model not found")
	}

	cek = 0

	db.Db.Find(&sparepart, "id_model = ?", id)

	for a := 0; a < len(sparepart); a++ {
		if sparepart[a].IdModel == id {
			if sparepart[a].Nama == tsp.SparePart {
				x = a
				break
			}
		}
		cek++
	}

	if cek == len(sparepart) {
		return errors.New("spare part not found")
	}

	tsp.TotalHarga = sparepart[x].Harga * tsp.Jumlah

	db.Db.First(&val, "id = ?", param)

	sparepart[x].Stok += val.Jumlah

	if tsp.Jumlah > sparepart[x].Stok {
		return errors.New("spare parts stock is running out")
	}

	sparepart[x].Stok -= tsp.Jumlah
	db.Db.Where("id = ?", sparepart[x].ID).Updates(&sparepart[x])
	db.Db.Where("id = ?", param).Updates(&tsp)
	return nil
}

func (db dbRepos) DeleteTsp(id int64) error {
	err := db.Db.Where("id = ?", id).Delete(&user.TransaksiSparepart{}).Error

	if err != nil {
		return err
	}
	return nil
}
