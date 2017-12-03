package controllers

import (
	"github.com/qiangxue/fasthttp-routing"
	"encoding/json"
	m "../models"
	"../dao"
	"github.com/valyala/fasthttp"
)

func CreateUser(ctx *routing.Context) error  {
	body := new(m.User)
	err := json.Unmarshal(ctx.PostBody(), body)
	if err != nil {
		return err
	}
	nickname := ctx.Param("nickname")
	body.Nickname = nickname
	var conflictUsers []m.User
	dao.GetUserByNickOrEmail(nickname, body.Email, &conflictUsers)
	if len(conflictUsers) == 0 {
		dao.CreateUser(body) // maybe check err
		myJSON(ctx.RequestCtx, fasthttp.StatusCreated, body)
		return nil
	} else {
		myJSON(ctx.RequestCtx, fasthttp.StatusConflict, conflictUsers)
		return nil
	}
}

func GetUser(ctx *routing.Context) error {
	nickname := ctx.Param("nickname")
	userExists := dao.GetUserByNickname(nickname) //Check if user exists
	if userExists == nil {
		myJSON(ctx.RequestCtx, fasthttp.StatusNotFound, nil) //User doesn't exist
		return nil
	}
	myJSON(ctx.RequestCtx, fasthttp.StatusOK, userExists)
	return nil
}


