package handlers

import (
	"net/http"
	"project_bengkel/bengkel"
	"project_bengkel/middleware"

	"github.com/gin-gonic/gin"
)

func NewHandlersTrans(us bengkel.UsecaseTrans, r *gin.RouterGroup) {
	uc := useTrans{
		uscase: us,
	}

	v2 := r.Group("Transaksi")

	v3 := v2.Group("Servicess")
	v3.GET("", middleware.Token(), uc.GetAllTserv)
	v3.POST("", uc.CreateTserv)
	v3.PUT(":id", uc.UpdateTserv)
	v3.DELETE("", uc.DeleteTserv)

	v4 := v2.Group("Spareparts")
	v4.GET("", middleware.Token(), uc.GetAllTsp)
	v4.POST("", uc.CreateTsp)
	v4.PUT(":id", uc.UpdateTsp)
	v4.DELETE("", uc.DeleteTsp)
}

type useTrans struct {
	uscase bengkel.UsecaseTrans
}

// Transaksi Service
func (usTs useTrans) GetAllTserv(ctx *gin.Context) {
	result, err := usTs.uscase.GetAllTserv(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (usTs useTrans) CreateTserv(ctx *gin.Context) {
	err := usTs.uscase.CreateTserv(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Post Data"})
}

func (usTs useTrans) UpdateTserv(ctx *gin.Context) {
	err := usTs.uscase.UpdateTserv(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Update Data"})
}

func (usTs useTrans) DeleteTserv(ctx *gin.Context) {
	err := usTs.uscase.DeleteTserv(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Delete Data"})
}

// Transaksi Sparepart
func (usTs useTrans) GetAllTsp(ctx *gin.Context) {
	result, err := usTs.uscase.GetAllTsp(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (usTs useTrans) CreateTsp(ctx *gin.Context) {
	err := usTs.uscase.CreateTsp(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Post Data"})
}

func (usTs useTrans) UpdateTsp(ctx *gin.Context) {
	err := usTs.uscase.UpdateTsp(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Update Data"})
}

func (usTs useTrans) DeleteTsp(ctx *gin.Context) {
	err := usTs.uscase.DeleteTsp(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Succes Delete Data"})
}
