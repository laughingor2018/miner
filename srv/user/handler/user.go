package handler

import (
	"context"
	"crypto/md5"
	"fmt"
	"time"

	"github.com/micro/go-log"
	user "github.com/laughingor2018/miner/srv/user/proto/user"
)

type User struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) Login(ctx context.Context, req *user.RequestLogin, rsp *user.ResponseLogin) error {
	log.Log("Received User.Login request")
	rsp.Code = 0

	if rsp.Code == 0 {
		rsp.Description = "success"

		m := md5.New()
		s := m.Sum([]byte(time.Now().String()))
		token := fmt.Sprintf("%x", s)

		rsp.Data = &user.Data{
			Uid:req.Uid,
			Account:req.Account,
			Token:token,
		}

	} else {
		rsp.Description = "login failed"
		rsp.Data = &user.Data{
			Uid:req.Uid,
			Account:req.Account,
		}
	}

	return nil
}


func (e *User) Register(ctx context.Context, req *user.RequestRegister, rsp *user.ResponseRegister) error {
	log.Log("Received User.Register request")
	rsp.Code = 0

	if rsp.Code == 0 {
		rsp.Description = "success"

		rsp.Data = &user.Data{
			Uid:1,
			Account:req.Account,
		}

	} else {
		rsp.Description = "register failed"
		rsp.Data = &user.Data{
			Uid:1,
			Account:req.Account,
		}
	}

	return nil
}


