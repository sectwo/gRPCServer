package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-server/config"
	"grpc-server/gRPC/paseto"
	auth "grpc-server/gRPC/proto"
)

type GRPCClient struct {
	client      *grpc.ClientConn
	authClient  auth.AuthServiceClient
	pasetoMaker *paseto.PasetoMaker
}

func NewGRPCClient(cfg *config.Config) (*GRPCClient, error) {
	c := new(GRPCClient)

	if client, err := grpc.Dial(cfg.GPRC.URL, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return nil, err
	} else {
		c.client = client
		c.authClient = auth.NewAuthServiceClient(client)
		c.pasetoMaker = paseto.NewPasetoMaker(cfg)
	}

	return c, nil
}

func (g *GRPCClient) CreateAuth(address string) (*auth.AuthData, error) {

	return nil, nil
}

func (g *GRPCClient) VerifyAuth(token string) (*auth.VerifyTokenRes, error) {

	return nil, nil
}
