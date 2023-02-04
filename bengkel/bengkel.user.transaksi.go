package bengkel

import (
	"project_bengkel/entity/user"

	"github.com/gin-gonic/gin"
)

type RepoTrans interface {
	GetAllTsp() ([]user.TransaksiSparepart, error)
	CreateTsp(tsp user.TransaksiSparepart) error
	UpdateTsp(tsp user.TransaksiSparepart, param string) error
	DeleteTsp(id int64) error

	GetAllTserv() ([]user.TransaksiService, error)
	CreateTserv(tsp user.TransaksiService) error
	UpdateTserv(tsp user.TransaksiService, param string) error
	DeleteTserv(id int64) error
}

type UsecaseTrans interface {
	GetAllTsp(ctx *gin.Context) ([]user.TransaksiSparepart, error)
	CreateTsp(ctx *gin.Context) error
	UpdateTsp(ctx *gin.Context) error
	DeleteTsp(ctx *gin.Context) error

	GetAllTserv(ctx *gin.Context) ([]user.TransaksiService, error)
	CreateTserv(ctx *gin.Context) error
	UpdateTserv(ctx *gin.Context) error
	DeleteTserv(ctx *gin.Context) error
}
