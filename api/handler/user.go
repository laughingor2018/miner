package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-log"

	"github.com/laughingor2018/miner/api/client"
	user "github.com/laughingor2018/miner/srv/user/proto/user"
	api "github.com/micro/go-api/proto"
	"github.com/micro/go-micro/errors"
)

type User struct{}

// Example.Call is called by the API as /user/example/call with post body {"name": "foo"}
func (e *User) Login(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received User.Login request")

	// extract the client from the context
	userClient, ok := client.UserFromContext(ctx)
	if !ok {
		return errors.InternalServerError("ars.miner.api.user.user.login", "user client not found")
	}

	// make request
	response, err := userClient.Login(ctx, &user.RequestLogin{
		Account: extractValue(req.Post["account"]),
		Password: extractValue(req.Post["password"]),
	})

	if err != nil {
		return errors.InternalServerError("ars.miner.api.user.user.login", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}

func (e *User) Register(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received User.Register request")

	// extract the client from the context
	userClient, ok := client.UserFromContext(ctx)
	if !ok {
		return errors.InternalServerError("ars.miner.api.user.user.register", "user client not found")
	}

	// make request
	response, err := userClient.Register(ctx, &user.RequestRegister{
		Account: extractValue(req.Post["account"]),
		Password: extractValue(req.Post["password"]),
	})

	if err != nil {
		return errors.InternalServerError("ars.miner.api.user.user.register", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
