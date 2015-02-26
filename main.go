package main

import (
	_ "github.com/igordonshaw/inote/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	beego.Run()
}

