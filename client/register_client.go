package client

import (
	"context"

	"github.com/begenov/api-gateway/pb"
	"google.golang.org/grpc"
)

type RegisterServiceClient struct {
	Conn   *grpc.ClientConn
	Client pb.RegisterClient
}

func NewRegisterServiceClient(addr string) (*RegisterServiceClient, error) {
	// Создайте соединение с удаленным микросервисом
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	// Создайте клиентский объект
	client := pb.NewRegisterClient(conn)

	return &RegisterServiceClient{
		Conn:   conn,
		Client: client,
	}, nil
}

func (c *RegisterServiceClient) SignUp(ctx context.Context, req *pb.RequestRegister) (*pb.Response, error) {
	return c.Client.SignUp(ctx, req)
}
