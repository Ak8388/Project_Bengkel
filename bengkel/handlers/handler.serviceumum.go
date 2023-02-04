package handlers

import (
	"net/http"
	"project_bengkel/bengkel"

	"github.com/gin-gonic/gin"
)

func NewHandlersSU(use bengkel.UseSU, r *gin.RouterGroup) {
	us := &useSu{
		usec: use,
	}

	v2 := r.Group("Service Umum")
	v2.GET("", us.GetAllSU)
	v2.POST("", us.CreateSU)
	v2.PUT(":id", us.UpdateSU)
	v2.DELETE("", us.DeleteSU)
}

type useSu struct {
	usec bengkel.UseSU
}

func (us useSu) GetAllSU(ctx *gin.Context) {
	result, err := us.usec.GetAllSU(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (us useSu) CreateSU(ctx *gin.Context) {
	err := us.usec.CreateSU(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Post Data"})
}

func (us useSu) UpdateSU(ctx *gin.Context) {
	err := us.usec.UpdateSU(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Update Data"})

}

func (us useSu) DeleteSU(ctx *gin.Context) {
	err := us.usec.DeleteSU(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Delete Data"})
}
