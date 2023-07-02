package main

import (
	"context"
	"log"

	"github.com/begenov/api-gateway/client"
	"github.com/begenov/api-gateway/pb"
)

func main() {
	client, err := client.NewRegisterServiceClient("localhost:9090")
	if err != nil {
		log.Fatalf("Не удалось создать клиентскую структуру: %v", err)
	}
	defer client.Conn.Close()

	ctx := context.Background()

	request := &pb.RequestRegister{
		Email: "asf",
		Phone: "1342",
		Role:  "user",
	}

	response, err := client.SignUp(ctx, request)
	if err != nil {
		log.Fatalf("Ошибка при вызове метода SignUp: %v", err)
	}

	log.Printf("Ответ от сервера: %s", response.Message)
}
