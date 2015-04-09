package main

import (
	_ "github.com/igordonshaw/inote/routers"
	"github.com/astaxie/beego"
	"github.com/igordonshaw/inote/controllers"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.SetLogger("file", `{"filename":"./log/inote.log"}`)
	beego.Run()
}

