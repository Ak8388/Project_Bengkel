package repository

import (
	"errors"
	"project_bengkel/entity"
	"strconv"

	"gorm.io/gorm"
)

func NewRepoSP(db *gorm.DB) *repoSp {
	return &repoSp{
		Db: db,
	}

}

type repoSp struct {
	Db *gorm.DB
}

func (db repoSp) GetAllSP() ([]entity.SparePart, error) {
	var result []entity.SparePart

	err := db.Db.Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (db repoSp) CreateSP(sp entity.SparePart) error {
	var val []entity.SparePart

	db.Db.Find(&val, "nama = ?", sp.Nama)

	for a := 0; a < len(val); a++ {
		if val[a].Nama == sp.Nama {
			if val[a].IdModel == sp.IdModel {
				return errors.New("Spare Part Already Exsit")
			}
		}
	}

	err := db.Db.Create(&sp).Error

	if err != nil {
		return err
	}

	return nil
}

func (db repoSp) UpdateSP(sp entity.SparePart, param string) error {
	var val []entity.SparePart

	db.Db.Find(&val, "nama = ?", sp.Nama)

	for a := 0; a < len(val); a++ {
		if val[a].Nama == sp.Nama {
			if val[a].IdModel == sp.IdModel {
				id, _ := strconv.ParseUint(param, 10, 64)
				if uint64(val[a].ID) != id {
					return errors.New("Spare Part Already Exsit")
				}
			}
		}
	}

	err := db.Db.Where("id = ?", param).Updates(&sp).Error

	if err != nil {
		return err
	}

	return nil
}

func (db repoSp) DeleteSP(id int64) error {
	err := db.Db.Where("id = ?", id).Delete(&entity.SparePart{}).Error
	if err != nil {
		return err
	}

	return nil
}
