package repository

import (
	"errors"
	"project_bengkel/entity"

	"gorm.io/gorm"
)

func NewRepoKat(db *gorm.DB) *dbkat {
	return &dbkat{
		Db: db,
	}
}

type dbkat struct {
	Db *gorm.DB
}

func (db dbkat) GetAllKat() ([]entity.Kategori, error) {
	var kategori []entity.Kategori
	db.Db.Find(&kategori)

	return kategori, nil
}

func (db dbkat) CreateKat(kat entity.Kategori) error {
	var val entity.Kategori
	db.Db.First(&val, "ka_name = ?", kat.KaName)

	if val.KaName != "" {
		return errors.New("Kategoris Name Already Exsit")
	}

	err := db.Db.Create(&kat).Error

	if err != nil {
		return err
	}

	return nil
}

func (db dbkat) UpdateKat(kat entity.Kategori, param string) error {
	var val entity.Kategori
	db.Db.First(&val, "name = ?", kat.KaName)

	if val.KaName != "" {
		return errors.New("Kategoris Name Already Exsit")
	}

	err := db.Db.Where("id = ?", param).Updates(&kat).Error

	if err != nil {
		return err
	}

	return nil
}

func (db dbkat) DeleteKat(param int64) error {
	var val entity.Kategori

	err := db.Db.First(&val, "id = ?", param).Error

	if err != nil {
		return errors.New("Data Not Found")
	}

	err = db.Db.Where("id_kat = ?", param).Delete(&entity.SparePart{}).Error

	if err != nil {
		return err
	}

	err = db.Db.Where("id = ?", param).Delete(&val).Error

	if err != nil {
		return err
	}

	return nil
}
