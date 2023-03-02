package delivery

import (
	"e-wallet/config"
	"e-wallet/delivery/controller"
	"e-wallet/manager"
	"fmt"

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
	controller.NewTransactionController(a.engine, a.useCaseManager.TransactionUseCase())
	fmt.Println("Eror mulai ==================================")
	controller.NewUserController(a.engine, a.useCaseManager.UserUseCase())
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run(":8082")
	if err != nil {
		panic(err.Error())
	}
}
