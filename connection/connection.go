package connection

import (
	"fmt"
	"log"
	"os"
	"project_bengkel/entity"
	"project_bengkel/entity/user"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Fatal("Can't Load Env")
	}

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbRoot := os.Getenv("DB_ROOT")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", dbUser, dbPass, dbName, dbRoot)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Can't Connect To Database")
	}

	db.AutoMigrate(&entity.MerkKendaraan{}, &entity.Service{}, &entity.ModelKendaraan{}, &entity.ServiceBerkala{}, &entity.ServiceUmum{}, &entity.Kategori{}, &entity.SparePart{}, &user.TransaksiService{}, &user.TransaksiSparepart{})
	db.AutoMigrate(&user.Regist{})
	return db
}
