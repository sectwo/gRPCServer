package paseto

import (
	"github.com/o1egl/paseto"
	"grpc-server/config"
)

type PasetoMaker struct {
	Pt  *paseto.V2
	Key []byte
}

func NewPasetoMaker(cfg config.Config) *PasetoMaker {

	return &PasetoMaker{
		Pt:  paseto.NewV2(),
		Key: []byte(cfg.Paseto.Key),
	}
}

func (m *PasetoMaker) CreateNewToken() (string, error) {
	return "", nil
}

func (m *PasetoMaker) VerifyToken(token string) error {
	return nil
}
