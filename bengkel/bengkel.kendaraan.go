package bengkel

import (
	"project_bengkel/entity"

	"github.com/gin-gonic/gin"
)

type UscaseMerk interface {
	GetAllMerk(ctx *gin.Context) ([]entity.MerkKendaraan, error)
	CreateMerk(ctx *gin.Context) error
	UpdateMerk(ctx *gin.Context) error
	DeleteMerk(ctx *gin.Context) error
}

type RepoMerk interface {
	GetAllMerk() ([]entity.MerkKendaraan, error)
	CreateMerk(data entity.MerkKendaraan) error
	UpdateMerk(data entity.MerkKendaraan, param string) error
	DeleteMerk(param int64) error
}
