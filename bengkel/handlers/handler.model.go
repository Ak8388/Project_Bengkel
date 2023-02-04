package handlers

import (
	"net/http"
	"project_bengkel/bengkel"

	"github.com/gin-gonic/gin"
)

func NewHandlersModel(model bengkel.UsecaseModelKendaraan, r *gin.RouterGroup) {
	md := useModel{
		usecase: model,
	}

	v2 := r.Group("Model")
	v2.GET(":model", md.GetByNamaMerk)
	v2.POST("", md.CreateModel)
	v2.PUT(":id", md.UpdateModel)
	v2.DELETE("", md.DeleteModel)
}

type useModel struct {
	usecase bengkel.UsecaseModelKendaraan
}

func (us useModel) GetByNamaMerk(ctx *gin.Context) {
	result, err := us.usecase.GetByNamaMerk(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Errors": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (us useModel) CreateModel(ctx *gin.Context) {
	err := us.usecase.CreateModel(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Errors": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusAccepted, gin.H{"Message": "Succes Post Data"})
}

func (us useModel) UpdateModel(ctx *gin.Context) {
	err := us.usecase.UpdateModel(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Errors": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusAccepted, gin.H{"Message": "Succes Update Data"})
}

func (us useModel) DeleteModel(ctx *gin.Context) {
	err := us.usecase.DeleteModel(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Errors": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusAccepted, gin.H{"Message": "Succes Delete Data"})
}
