package bengkel

import (
	"project_bengkel/entity"

	"github.com/gin-gonic/gin"
)

type RepoModelKendaraan interface {
	GetByNamaMerk(param string) ([]entity.ModelKendaraan, error)
	CreateModel(model entity.ModelKendaraan) error
	UpdateModel(model entity.ModelKendaraan, param string) error
	DeleteModel(id int64) error
}

type UsecaseModelKendaraan interface {
	GetByNamaMerk(ctx *gin.Context) ([]entity.ModelKendaraan, error)
	CreateModel(ctx *gin.Context) error
	UpdateModel(ctx *gin.Context) error
	DeleteModel(ctx *gin.Context) error
}
