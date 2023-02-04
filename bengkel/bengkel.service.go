package bengkel

import (
	"project_bengkel/entity"

	"github.com/gin-gonic/gin"
)

type UscaseService interface {
	GetAllService(ctx *gin.Context) ([]entity.Service, error)
	CreateService(ctx *gin.Context) error
	UpdateService(ctx *gin.Context) error
	DeleteService(ctx *gin.Context) error
}

type RepoService interface {
	GetAllService() ([]entity.Service, error)
	CreateService(data entity.Service) error
	UpdateService(data entity.Service, param string) error
	DeleteService(param int64) error
}
