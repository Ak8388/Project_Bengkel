package handlers

import (
	"net/http"
	"project_bengkel/bengkel"

	"github.com/gin-gonic/gin"
)

func NewHandlersKat(use bengkel.UseKat, r *gin.RouterGroup) {
	uskat := usekat{
		us: use,
	}

	v2 := r.Group("Kategori")
	v2.GET("", uskat.GetAllKat)
	v2.POST("", uskat.CreateKat)
	v2.PUT(":id", uskat.UpdateKat)
	v2.DELETE("", uskat.DeleteKat)
}

type usekat struct {
	us bengkel.UseKat
}

func (use usekat) GetAllKat(ctx *gin.Context) {
	result, err := use.us.GetAllKat(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (use usekat) CreateKat(ctx *gin.Context) {
	err := use.us.CreateKat(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Errors": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Post Data"})
}

func (use usekat) UpdateKat(ctx *gin.Context) {
	err := use.us.UpdateKat(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Errors": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Update Data"})

}

func (use usekat) DeleteKat(ctx *gin.Context) {
	err := use.us.DeleteKat(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Errors": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Delete Data"})
}
