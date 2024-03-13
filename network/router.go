package network

import (
	"github.com/gin-gonic/gin"
	"grpc-server/config"
	"grpc-server/service"
)

// repo 와 router 연결 에정

type Network struct {
	cfg *config.Config

	service *service.Service

	engin *gin.Engine
}

func NewNetwork(cfg *config.Config, service *service.Service) (*Network, error) {
	r := &Network{
		cfg:     cfg,
		service: service,
		engin:   gin.New(),
	}

	return r, nil
}

func (n *Network) StartServer() {
	_ = n.engin.Run(":8080")
}
