package authrepo

import (
	"project_bengkel/entity/user"

	"gorm.io/gorm"
)

type Db struct {
	db *gorm.DB
}

func NewReposAuth(database *gorm.DB) Db {
	return Db{
		db: database,
	}
}

func (database Db) Register(user user.Regist) error {
	err := database.db.Create(&user).Error

	if err != nil {
		return err
	}

	return nil
}
func (database Db) Login(username string) (*user.Regist, error) {
	var data user.Regist

	if err := database.db.First(&data, "username=?", username).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
func (database Db) Logout() error {
	return nil
}
