package repository

import (
	"errors"
	"project_bengkel/entity"

	"gorm.io/gorm"
)

type db struct {
	Db *gorm.DB
}

func NewRepo(Db *gorm.DB) *db {
	return &db{
		Db: Db,
	}
}

func (database db) GetAllService() ([]entity.Service, error) {
	var data []entity.Service
	database.Db.Find(&data)

	return data, nil
}

func (database db) CreateService(data entity.Service) error {
	var val entity.Service
	database.Db.First(&val, "type_service = ?", data.TypeService)

	if val.TypeService != "" {
		return errors.New("service already exsit")
	}

	err := database.Db.Create(&data).Error

	if err != nil {
		return err
	}

	return nil
}

func (database db) UpdateService(data entity.Service, param string) error {
	var val entity.Service
	database.Db.First(&val, "type_service = ?", data.TypeService)

	if val.TypeService != "" {
		return errors.New("service already exsit")
	}

	err := database.Db.Updates(&data).Error

	if err != nil {
		return err
	}

	return nil
}

func (database db) DeleteService(param int64) error {
	var data entity.Service
	err := database.Db.First(&data, "id = ?", param).Error

	if err != nil {
		return err
	}

	err = database.Db.Where("id_service=?", param).Delete(&entity.ServiceBerkala{}).Error

	if err != nil {
		return err
	}

	err = database.Db.Where("id_service=?", param).Delete(&entity.ServiceUmum{}).Error

	if err != nil {
		return err
	}

	err = database.Db.Where("id=?", param).Delete(&data).Error

	if err != nil {
		return err
	}

	return nil
}
