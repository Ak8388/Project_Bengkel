package bengkel

import (
	"project_bengkel/entity"

	"github.com/gin-gonic/gin"
)

type RepoKat interface {
	GetAllKat() ([]entity.Kategori, error)
	CreateKat(kat entity.Kategori) error
	UpdateKat(kat entity.Kategori, param string) error
	DeleteKat(param int64) error
}

type UseKat interface {
	GetAllKat(ctx *gin.Context) ([]entity.Kategori, error)
	CreateKat(ctx *gin.Context) error
	UpdateKat(ctx *gin.Context) error
	DeleteKat(ctx *gin.Context) error
}
