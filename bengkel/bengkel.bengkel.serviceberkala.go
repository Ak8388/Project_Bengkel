package bengkel

import (
	"project_bengkel/entity"

	"github.com/gin-gonic/gin"
)

type RepoSB interface {
	GetAllSB() ([]entity.ServiceBerkala, error)
	CreateSB(sb entity.ServiceBerkala) error
	UpdateSB(sb entity.ServiceBerkala, param string) error
	DeleteSB(param int64) error
}

type UseSB interface {
	GetAllSB(ctx *gin.Context) ([]entity.ServiceBerkala, error)
	CreateSB(ctx *gin.Context) error
	UpdateSB(ctx *gin.Context) error
	DeleteSB(ctx *gin.Context) error
}
