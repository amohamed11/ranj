package server

import "github.com/beego/beego/v2/server/web"

func StartServer(host string, port string) {
	ctrl := &MainController{}

	web.Router("/", ctrl, "get:Index")
	web.Run(host + ":" + port)
}

type MainController struct {
	web.Controller
}

// address: http://localhost:8080/ GET
func (ctrl *MainController) Index() {
	ctrl.Ctx.WriteString("hello, world")
}
