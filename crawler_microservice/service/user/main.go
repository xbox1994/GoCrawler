package main

import (
	"GoCrawler/crawler_microservice/service/user/proto"
	"context"
	"errors"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

type User struct {
	TokenService Authable
}

func (u *User) Login(ctx context.Context, req *user.User, rsp *user.Token) error {
	if req.Username != req.Password {
		return errors.New("wrong username or password")
	}

	token, e := u.TokenService.Encode(req)
	if e != nil {
		return e
	}

	rsp.Token = token
	return nil
}

func (u *User) ValidateToken(ctx context.Context, req *user.Token, rsp *user.Token) error {
	decode, e := u.TokenService.Decode(req.Token)
	if e != nil {
		return e
	}

	if decode.User.Id == "" {
		rsp.Error = &user.Error{
			Code:    -1,
			Message: "xx",
		}
		rsp.Valid = false
		return errors.New("invalid user")
	}

	rsp.Valid = true
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),

		// Setup some flags. Specify --run_client to run the client

		// Add runtime flags
		// We could do this below too
		micro.Flags(cli.BoolFlag{
			Name:  "run_client",
			Usage: "Launch the client",
		}),
	)

	service.Init()

	user.RegisterUserServiceHandler(service.Server(), new(User))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
