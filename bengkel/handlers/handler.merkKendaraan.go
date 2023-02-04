package handlers

import (
	"net/http"
	"project_bengkel/bengkel"

	"github.com/gin-gonic/gin"
)

func NewHandlersMerk(b bengkel.UscaseMerk, r *gin.RouterGroup) {
	use := &usecaseM{
		us: b,
	}

	v2 := r.Group("Merk Kendaraan")
	v2.GET("", use.GetAllMerk)
	v2.POST("", use.CreateMerk)
	v2.PUT(":id", use.UpdateMerk)
	v2.DELETE("", use.DeleteMerk)
}

type usecaseM struct {
	us bengkel.UscaseMerk
}

func (us usecaseM) GetAllMerk(ctx *gin.Context) {
	result, err := us.us.GetAllMerk(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (us usecaseM) CreateMerk(ctx *gin.Context) {
	err := us.us.CreateMerk(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Errors": err.Error(),
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Post Data"})
}

func (us usecaseM) UpdateMerk(ctx *gin.Context) {
	err := us.us.UpdateMerk(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Update Data"})
}

func (us usecaseM) DeleteMerk(ctx *gin.Context) {
	err := us.us.DeleteMerk(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Delete Data"})
}
