package cmd

import (
	"grpc-server/config"
	"grpc-server/network"
	"grpc-server/repository"
	"grpc-server/service"
)

// network, repository 에 대한 객체값과 기본적인 세팅 진행

type App struct {
	cfg *config.Config

	service    *service.Service
	repository *repository.Repository
	network    *network.Network
}

func NewApp(cfg *config.Config) {
	a := &App{cfg: cfg}

	var err error

	if a.repository, err = repository.NewRepository(cfg); err != nil {
		panic(err)
	} else if a.service, err = service.NewService(cfg, a.repository); err != nil {
		panic(err)
	} else if a.network, err = network.NewNetwork(cfg, a.service); err != nil {
		panic(err)
	} else {
		// TODO -> start server
		a.network.StartServer()

	}
}
