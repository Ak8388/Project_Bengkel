package handlers

import (
	"net/http"
	"project_bengkel/bengkel"

	"github.com/gin-gonic/gin"
)

func NewHandlersService(b bengkel.UscaseService, r *gin.RouterGroup) {
	use := &usecase{
		us: b,
	}

	// Service CRUD
	v2 := r.Group("Service")
	v2.GET("", use.GetAllService)
	v2.POST("", use.CreateService)
	v2.PUT(":id", use.UpdateService)
	v2.DELETE("", use.DeleteService)
}

type usecase struct {
	us bengkel.UscaseService
}

func (us usecase) GetAllService(ctx *gin.Context) {
	result, err := us.us.GetAllService(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (us usecase) CreateService(ctx *gin.Context) {
	err := us.us.CreateService(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Post Data"})
}

func (us usecase) UpdateService(ctx *gin.Context) {
	err := us.us.UpdateService(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Update Data"})
}

func (us usecase) DeleteService(ctx *gin.Context) {
	err := us.us.DeleteService(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Delete Data"})
}
