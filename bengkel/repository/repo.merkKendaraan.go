package repository

import (
	"errors"
	"project_bengkel/entity"

	"gorm.io/gorm"
)

type dbM struct {
	Db *gorm.DB
}

func NewRepoMerk(Db *gorm.DB) *dbM {
	return &dbM{
		Db: Db,
	}
}

func (database dbM) GetAllMerk() ([]entity.MerkKendaraan, error) {
	var data []entity.MerkKendaraan
	database.Db.Find(&data)

	return data, nil
}

func (database dbM) CreateMerk(data entity.MerkKendaraan) error {
	err := database.Db.First(&data, "merk = ?", data.Merk).Error

	if err == nil {
		return errors.New("merk already exsit")
	}

	err = database.Db.Create(&data).Error

	if err != nil {
		return err
	}

	return nil
}

func (database dbM) UpdateMerk(data entity.MerkKendaraan, param string) error {
	err := database.Db.First(&data, "merk = ?", data.Merk).Error

	if err == nil {
		return errors.New("merk already exsit")
	}

	database.Db.Where("id=?", param).Updates(&data)

	return nil
}

func (database dbM) DeleteMerk(param int64) error {
	var data entity.MerkKendaraan
	err := database.Db.First(&data, "id = ?", param).Error

	if err != nil {
		return err
	}

	err = database.Db.Where("id = ?", param).Delete(&data).Error

	if err != nil {
		return err
	}

	return nil
}
