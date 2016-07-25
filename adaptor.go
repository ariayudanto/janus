package main

import (
	"net/http"

	"github.com/kataras/iris"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func ToHandler(handler interface{}) iris.Handler {
	h := fasthttpadaptor.NewFastHTTPHandlerFunc(handler.(http.Handler).ServeHTTP)
	return ToHandlerFastHTTP(h)
}

func ToHandlerFastHTTP(h fasthttp.RequestHandler) iris.Handler {
	return iris.HandlerFunc((func(ctx *iris.Context) {
		h(ctx.RequestCtx)
		ctx.Next()
	}))
}
