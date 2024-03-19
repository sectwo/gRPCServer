package server

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-server/config"
	"grpc-server/gRPC/paseto"
	auth "grpc-server/gRPC/proto"
	"log"
	"net"
	"time"
)

// 3T 아키텍쳐 구조 대신 맵을 사용하여 샘플 구성
type GRPCServer struct {
	auth.AuthServiceServer
	pasetoMaker    *paseto.PasetoMaker
	tokenVerfiyMap map[string]*auth.AuthData
}

func NewGRPCServer(cfg *config.Config) error {
	if lis, err := net.Listen("tcp", cfg.GPRC.URL); err != nil {
		return err
	} else {

		server := grpc.NewServer([]grpc.ServerOption{}...)

		// AuthServiceServer
		// RegisterAuthServiceServer
		auth.RegisterAuthServiceServer(server, &GRPCServer{
			pasetoMaker:    paseto.NewPasetoMaker(cfg),
			tokenVerfiyMap: make(map[string]*auth.AuthData),
		})

		reflection.Register(server)

		go func() {
			log.Println("Server Start")
			if err = server.Serve(lis); err != nil {
				panic(err)
			}
		}()

	}
}

func (s *GRPCServer) CreateAuth(_ context.Context, req *auth.CreateTokenReq) (*auth.CreateTokenRes, error) {
	data := req.Auth
	token := data.Token
	s.tokenVerfiyMap[token] = data

	return &auth.CreateTokenRes{Auth: data}, nil
}

func (s *GRPCServer) VerifyAuth(_ context.Context, req *auth.VerifyTokenReq) (*auth.VerifyTokenRes, error) {
	token := req.Token

	res := &auth.VerifyTokenRes{V: &auth.Verify{
		Auth: nil,
	}}

	if authData, ok := s.tokenVerfiyMap[token]; !ok {
		res.V.Status = auth.ResponseType_FAILED
	} else if authData.ExpireDate < time.Now().Unix() {
		delete(s.tokenVerfiyMap, token)
		res.V.Status = auth.ResponseType_EXPIRED_DATE
	} else {
		res.V.Status = auth.ResponseType_SUCCESS
	}

	return res, nil
}

func (s *GRPCServer) mustEmbedUnimplementedAuthServiceServer() {
	//TODO implement me
	panic("implement me")
}
