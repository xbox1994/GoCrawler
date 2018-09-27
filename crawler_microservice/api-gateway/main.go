package main

import (
	"GoCrawler/crawler_microservice/api-gateway/proto"
	"GoCrawler/crawler_microservice/service/user/proto"
	"context"
	"encoding/json"
	"errors"
	api "github.com/micro/go-api/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"log"
	"strings"
)

type Gateway struct{}

func (e *Gateway) Login(ctx context.Context, req *api.Request, rsp *api.Response) error {
	authClient := user.NewUserService("go.micro.srv.user", client.DefaultClient)
	authResp, err := authClient.Login(context.Background(), &user.User{
		Id:       "id",
		Username: "username",
		Password: "username",
	})

	if err != nil {
		return err
	}

	log.Println("Login Resp:", authResp)
	b, _ := json.Marshal(map[string]string{
		"token": authResp.Token,
	})
	rsp.Body = string(b)
	return err
}

func (e *Gateway) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Gateway.Call request")

	// parse values from the get request
	name, ok := req.Get["name"]

	if !ok || len(name.Values) == 0 {
		return errors.New("go.micro.api.example no content")
	}

	// set response status
	rsp.StatusCode = 200

	// respond with some json
	b, _ := json.Marshal(map[string]string{
		"message": "got your request " + strings.Join(name.Values, " "),
	})

	// set json body
	rsp.Body = string(b)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.gateway"),
		micro.WrapHandler(AuthWrapper),
	)

	service.Init()

	api_gateway.RegisterGatewayHandler(service.Server(), new(Gateway))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		if req.Service() == "go.micro.api.gateway" && req.Method() == "Gateway.Login" {
			err := fn(ctx, req, resp)
			return err
		}

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
