package repository

import (
	"errors"
	"project_bengkel/entity"

	"gorm.io/gorm"
)

func NewRepoMod(db *gorm.DB) *database {
	return &database{
		Db: db,
	}
}

type database struct {
	Db *gorm.DB
}

func (db database) GetByNamaMerk(param string) ([]entity.ModelKendaraan, error) {
	var model []entity.ModelKendaraan
	var merk entity.MerkKendaraan
	err := db.Db.First(&entity.MerkKendaraan{}, "merk = ?", param).Error

	if err != nil {
		return nil, err
	}

	db.Db.Select("id").Find(&merk, "merk = ?", param)

	db.Db.Where("id_merk = ?", merk.ID).Find(&model)
	return model, nil
}

func (db database) CreateModel(model entity.ModelKendaraan) error {
	var mod []entity.ModelKendaraan

	db.Db.Find(&mod, "tahun = ?", model.Tahun)

	for a := 0; a < len(mod); a++ {
		if mod[a].Model == model.Model {
			if mod[a].Tahun == model.Tahun {
				return errors.New("model already exsit")
			}
		}
	}

	err := db.Db.Create(&model).Error

	if err != nil {
		return err
	}

	return nil
}

func (db database) UpdateModel(model entity.ModelKendaraan, param string) error {
	var mod []entity.ModelKendaraan

	db.Db.Find(&mod, "tahun = ?", model.Tahun)

	for a := 0; a < len(mod); a++ {
		if mod[a].Model == model.Model {
			if mod[a].Tahun == model.Tahun {
				return errors.New("model already exsit")
			}
		}
	}

	err := db.Db.Where("id = ? ", param).Updates(&model).Error

	if err != nil {
		return err
	}

	return nil
}

func (db database) DeleteModel(id int64) error {
	err := db.Db.Where("id_model = ?", id).Delete(&entity.SparePart{}).Error

	if err != nil {
		return err
	}

	err = db.Db.Where("id = ?", id).Delete(&entity.ModelKendaraan{}).Error

	if err != nil {
		return err
	}

	return nil
}
