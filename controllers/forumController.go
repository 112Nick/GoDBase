package controllers

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/qiangxue/fasthttp-routing"
	m "../models"
	"../dao"
	"github.com/fasthttp-contrib/render"
	"fmt"
)

var response = render.New()

func CreateForum(ctx *routing.Context) error {
	body := new(m.Forum)
	err := json.Unmarshal(ctx.PostBody(), body)
	if err != nil {
		return err
	}
	userExists:= dao.GetUserByNickname(body.User) //Check if user exists
	if userExists == nil {
		myJSON(ctx.RequestCtx, fasthttp.StatusNotFound, nil) //User doesn't exist
		return nil
	}
	forum:= dao.GetForumBySlug(body.Slug)
	if forum != nil {
		myJSON(ctx.RequestCtx, fasthttp.StatusConflict, forum)
		return nil
	}
 	body.User = userExists.Nickname
	err = dao.CreateForum(body) //Trying to create forum, may cause conflict
	fmt.Println(err)
	myJSON(ctx.RequestCtx, fasthttp.StatusCreated, body) //StatusCreated
	return nil
}