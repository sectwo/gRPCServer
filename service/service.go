package service

import (
	"grpc-server/config"
	"grpc-server/repository"
)

// repo 와 router 연결 에정

type Service struct {
	cfg *config.Config

	repository *repository.Repository
}

func NewService(cfg *config.Config, repository *repository.Repository) (*Service, error) {
	r := &Service{
		cfg:        cfg,
		repository: repository,
	}

	return r, nil
}
