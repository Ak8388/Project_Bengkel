package bengkel

import (
	"project_bengkel/entity"

	"github.com/gin-gonic/gin"
)

type RepoSU interface {
	GetAllSU() ([]entity.ServiceUmum, error)
	CreateSU(su entity.ServiceUmum) error
	UpdateSU(su entity.ServiceUmum, param string) error
	DeleteSU(param int64) error
}

type UseSU interface {
	GetAllSU(ctx *gin.Context) ([]entity.ServiceUmum, error)
	CreateSU(ctx *gin.Context) error
	UpdateSU(ctx *gin.Context) error
	DeleteSU(ctx *gin.Context) error
}
