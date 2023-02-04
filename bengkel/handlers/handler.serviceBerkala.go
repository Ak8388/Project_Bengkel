package handlers

import (
	"net/http"
	"project_bengkel/bengkel"

	"github.com/gin-gonic/gin"
)

func NewHandlersSB(use bengkel.UseSB, r *gin.RouterGroup) {
	ussb := &usec{
		uscase: use,
	}

	v2 := r.Group("Service Berkala")
	v2.GET("", ussb.GetAllSB)
	v2.POST("", ussb.CreateSB)
	v2.PUT(":id", ussb.UpdateSB)
	v2.DELETE("", ussb.DeleteSB)
}

type usec struct {
	uscase bengkel.UseSB
}

func (us usec) GetAllSB(ctx *gin.Context) {
	result, err := us.uscase.GetAllSB(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Service Berkala": result})
}

func (us usec) CreateSB(ctx *gin.Context) {
	err := us.uscase.CreateSB(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Post Data"})
}

func (us usec) UpdateSB(ctx *gin.Context) {
	err := us.uscase.UpdateSB(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Update Data"})
}

func (us usec) DeleteSB(ctx *gin.Context) {
	err := us.uscase.DeleteSB(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Delete Data"})
}
