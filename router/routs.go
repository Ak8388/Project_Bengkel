package router

import (
	"project_bengkel/bengkel/handlers"
	"project_bengkel/bengkel/repository"
	"project_bengkel/bengkel/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routes struct {
	Db *gorm.DB
	R  *gin.Engine
}

func (r Routes) Routers() {
	v1 := r.R.Group("Api")
	//Service Routers
	repo := repository.NewRepo(r.Db)
	use := usecase.NewUsecaseService(repo)
	handlers.NewHandlersService(use, v1)

	//Merk Routers
	repoM := repository.NewRepoMerk(r.Db)
	useM := usecase.NewUsecaseMerk(repoM)
	handlers.NewHandlersMerk(useM, v1)

	//Model Routers
	repoMod := repository.NewRepoMod(r.Db)
	useMod := usecase.NewUsecaseModel(repoMod)
	handlers.NewHandlersModel(useMod, v1)

	//Service Berkala Routers
	repoSB := repository.NewRepoSB(r.Db)
	useSB := usecase.NewUsecaseSB(repoSB)
	handlers.NewHandlersSB(useSB, v1)

	//Service Umum Routers
	repoSU := repository.NewRepoSU(r.Db)
	useSU := usecase.NewUsecaseSU(repoSU)
	handlers.NewHandlersSU(useSU, v1)

	//Kategori Routers
	repoKat := repository.NewRepoKat(r.Db)
	useKat := usecase.NewUsecaseKat(repoKat)
	handlers.NewHandlersKat(useKat, v1)

	//Sparepart Routers
	repoSP := repository.NewRepoSP(r.Db)
	useSp := usecase.NewUsecaseSP(repoSP)
	handlers.NewHandlersSP(useSp, v1)

	//Transaksi
	RepoT := repository.NewRepoTrans(r.Db)
	UseT := usecase.NewUsecaseTrans(RepoT)
	handlers.NewHandlersTrans(UseT, v1)
}
