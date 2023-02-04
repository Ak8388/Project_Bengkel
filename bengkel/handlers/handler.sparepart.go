package handlers

import (
	"net/http"
	"project_bengkel/bengkel"

	"github.com/gin-gonic/gin"
)

func NewHandlersSP(us bengkel.UsecaseSparePart, r *gin.RouterGroup) {
	usSp := useSp{
		usecase: us,
	}

	v2 := r.Group("Sparepart")
	v2.GET("", usSp.GetAllSparePart)
	v2.POST("", usSp.CreateSparePart)
	v2.PUT(":id", usSp.UpdateSparePart)
	v2.DELETE("", usSp.DeleteSparePart)
}

type useSp struct {
	usecase bengkel.UsecaseSparePart
}

func (us useSp) GetAllSparePart(ctx *gin.Context) {
	result, err := us.usecase.GetAllSP(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (us useSp) CreateSparePart(ctx *gin.Context) {
	err := us.usecase.CreateSP(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Post Data"})
}

func (us useSp) UpdateSparePart(ctx *gin.Context) {
	err := us.usecase.UpdateSP(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Update Data"})
}

func (us useSp) DeleteSparePart(ctx *gin.Context) {
	err := us.usecase.DeleteSP(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Delete Data"})
}
