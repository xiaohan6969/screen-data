package router

import (
	"../controller/homepage"
	"../middleware/corsServer"
	"github.com/kataras/iris"
)

func Router() *iris.Application {
	app := iris.Default()

	app.WrapRouter(corsServer.Cors().ServeHTTP) // 跨域请求

	app.Handle("GET", "/", homepage.IndexHtml) //首页

	return app
}