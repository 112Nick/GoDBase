package controllers

import (
	"github.com/valyala/fasthttp"
)

func myJSON (ctx *fasthttp.RequestCtx, status int, v interface{}) {
	response.JSON(ctx, status, v)
}
