package repository

import (
	"errors"
	"project_bengkel/entity"

	"gorm.io/gorm"
)

func NewRepoSU(db *gorm.DB) *reposu {
	return &reposu{
		Db: db,
	}
}

type reposu struct {
	Db *gorm.DB
}

func (db reposu) GetAllSU() ([]entity.ServiceUmum, error) {
	var su []entity.ServiceUmum

	db.Db.Find(&su)

	return su, nil
}

func (db reposu) CreateSU(su entity.ServiceUmum) error {
	var val entity.ServiceUmum
	db.Db.First(&val, "service = ?", su.Service)

	if val.Service != "" {
		return errors.New("service already exsit")
	}

	err := db.Db.Create(&su).Error

	if err != nil {
		return err
	}

	return nil
}

func (db reposu) UpdateSU(su entity.ServiceUmum, param string) error {
	var val entity.ServiceUmum
	db.Db.First(&val, "service = ?", su.Service)

	if val.Service != "" {
		return errors.New("service already exsit")
	}

	err := db.Db.Create(&su).Error

	if err != nil {
		return err
	}

	return nil
}

func (db reposu) DeleteSU(param int64) error {
	var sb entity.ServiceBerkala

	err := db.Db.Where("id = ?", param).Delete(&sb).Error

	if err != nil {
		return err
	}

	return nil
}
