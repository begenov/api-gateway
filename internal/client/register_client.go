package client

import (
	"context"

	"github.com/begenov/register-service/pb"
	"google.golang.org/grpc"
)

type RegisterServiceClient struct {
	Conn   *grpc.ClientConn
	Client pb.RegisterClient
}

func NewRegisterServiceClient(addr string) (*RegisterServiceClient, error) {
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

func (c *RegisterServiceClient) SignIn(ctx context.Context, req *pb.RequestSignIn) (*pb.ResponseToken, error) {
	return c.Client.SignIn(ctx, req)
}

func (c *RegisterServiceClient) RefreshToken(ctx context.Context, req *pb.RequestToken) (*pb.ResponseToken, error) {
	return c.Client.RefreshToken(ctx, req)
}
