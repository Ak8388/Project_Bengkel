package repository

import (
	"errors"
	"project_bengkel/entity"

	"gorm.io/gorm"
)

func NewRepoSB(db *gorm.DB) *dbSb {
	return &dbSb{
		Db: db,
	}
}

type dbSb struct {
	Db *gorm.DB
}

func (db dbSb) GetAllSB() ([]entity.ServiceBerkala, error) {
	var sb []entity.ServiceBerkala

	db.Db.Find(&sb)

	return sb, nil
}

func (db dbSb) CreateSB(sb entity.ServiceBerkala) error {
	var val entity.ServiceBerkala
	db.Db.First(&val, "service = ?", sb.Service)

	if val.Service != "" {
		return errors.New("service already exsit")
	}

	err := db.Db.Create(&sb).Error

	if err != nil {
		return err
	}

	return nil
}

func (db dbSb) UpdateSB(sb entity.ServiceBerkala, param string) error {
	var val entity.ServiceBerkala
	db.Db.First(&val, "service = ?", sb.Service)

	if val.Service != "" {
		return errors.New("service already exsit")
	}

	err := db.Db.Create(&sb).Error

	if err != nil {
		return err
	}

	return nil
}

func (db dbSb) DeleteSB(param int64) error {
	var sb entity.ServiceBerkala

	err := db.Db.Where("id = ?", param).Delete(&sb).Error

	if err != nil {
		return err
	}

	return nil
}
