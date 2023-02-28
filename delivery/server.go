package delivery

import (
	"e-wallet/config"
	"e-wallet/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	engine         *gin.Engine
	useCaseManager manager.UseCaseManager
}

func Server() *appServer {
	ginEngine := gin.Default()
	config := config.NewConfig()
	infra := manager.NewInfraManager(config)
	repo := manager.NewRepositoryManager(infra)
	usecase := manager.NewUseCaseManager(repo)
	return &appServer{
		engine:         ginEngine,
		useCaseManager: usecase,
	}
}

func (a *appServer) initHandlers() {
	//masukan yang mau di jalankan handlernya dari package controller
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run()
	if err != nil {
		panic(err.Error())
	}
}
