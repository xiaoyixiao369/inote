package routers

import (
    "github.com/igordonshaw/inote/controllers"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/context"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &controllers.MainController{},"get:LoginPage")
    nsApi :=
        beego.NewNamespace("/i",
        beego.NSRouter("/posts/:id", &controllers.PostController{}, "get:One"),
    )
    beego.AddNamespace(nsApi)

    var checkUser = func(ctx *context.Context) {
        /*_, ok := ctx.Input.Session("uid").(int)
        if !ok && ctx.Request.RequestURI != "/login" {
            ctx.Redirect(302, "/login") }*/

    }
    beego.InsertFilter("/*",beego.BeforeRouter,checkUser)
}
