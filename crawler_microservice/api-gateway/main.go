package main

import (
	hello "GoCrawler/crawler_microservice/service/greeter/proto"
	"GoCrawler/crawler_microservice/service/user/proto"
	"context"
	"errors"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"log"
)

type Say struct{}

var (
	cl hello.GreeterService
)

func (s *Say) Hello(ctx context.Context, req *hello.HelloRequest, rsp *hello.HelloResponse) error {
	log.Print("Received Say.Hello API request")

	name := req.Name

	response, err := cl.Hello(context.TODO(), &hello.HelloRequest{
		Name: name,
	})

	if err != nil {
		return err
	}

	rsp.Greeting = response.Greeting

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.gateway"),
		micro.WrapHandler(AuthWrapper),
	)

	service.Init()

	cl = hello.NewGreeterService("go.micro.srv.greeter", client.DefaultClient)

	hello.RegisterGreeterHandler(service.Server(), new(Say))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {

		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		token := meta["Token"]

		authClient := user.NewUserService("go.micro.srv.user", client.DefaultClient)
		authResp, err := authClient.ValidateToken(context.Background(), &user.Token{
			Token: token,
		})
		log.Println("Auth Resp:", authResp)
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}
