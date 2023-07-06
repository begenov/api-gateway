package main

import (
	"fmt"

	"github.com/begenov/api-gateway/internal/app"
	"github.com/begenov/api-gateway/internal/config"
)

func main() {
	cfg := config.NewConfig()
	fmt.Println(cfg)
	if err := app.Run(cfg); err != nil {
		panic(err)
	}
}

/*
client, err := client.NewRegisterServiceClient("localhost:9090")
	if err != nil {
		log.Fatalf("Не удалось создать клиентскую структуру: %v", err)
	}
	defer client.Conn.Close()

	ctx := context.Background()

	// test, err := client.SignUp(ctx, &pb.RequestRegister{
	// 	Email:    "test1@gmail.com",
	// 	Password: "test",
	// 	Address:  "test",
	// 	Role:     "user",
	// 	Phone:    "7889456",
	// })
	// if err != nil {
	// 	log.Fatal(err, "sign-up")
	// }

	// fmt.Printf("test.Message: %v\n", test.Message)

	// request := &pb.RequestSignIn{
	// 	Email:    "test1@gmail.com",
	// 	Password: "test",
	// 	Role:     "user",
	// }

	// response, err := client.SignIn(ctx, request)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("Ответ от сервера: %s \t %s", response.AccessToken, response.RefreshToken)

	inp := &pb.RequestToken{
		RefreshToken: "$2a$10$tEHCmZEe3VznM7t2vJZ1r.aBJ1xkFBkxgYcHpMsnO.Cin1VzJUKQa",
		Role:         "user",
	}

	token, err := client.RefreshToken(ctx, inp)

	if err != nil {
		log.Fatal(err, "--token--")
	}

	fmt.Printf("token.AccessToken: %v\n", token.AccessToken)
	fmt.Printf("token.RefreshToken: %v\n", token.RefreshToken)
*/
