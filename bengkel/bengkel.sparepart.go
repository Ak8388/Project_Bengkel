package bengkel

import (
	"project_bengkel/entity"

	"github.com/gin-gonic/gin"
)

type RepoSparePart interface {
	GetAllSP() ([]entity.SparePart, error)
	CreateSP(sp entity.SparePart) error
	UpdateSP(sp entity.SparePart, param string) error
	DeleteSP(id int64) error
}

type UsecaseSparePart interface {
	GetAllSP(ctx *gin.Context) ([]entity.SparePart, error)
	CreateSP(ctx *gin.Context) error
	UpdateSP(ctx *gin.Context) error
	DeleteSP(ctx *gin.Context) error
}
