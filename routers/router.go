package routers

import (
    "github.com/igordonshaw/inote/controllers"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/context"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &controllers.MainController{}, "get:LoginPage")
    beego.Router("/author", &controllers.UserControlelr{}, "get:Author")
    nsApi :=
        beego.NewNamespace("/i",
        beego.NSRouter("/posts", &controllers.PostController{}, "get:Posts"),
        beego.NSRouter("/posts/list", &controllers.PostController{}, "get:ListPosts"),
        beego.NSRouter("/posts/:id", &controllers.PostController{}, "get:OnePost"),
    )
    beego.AddNamespace(nsApi)

    nsAdmin :=
       beego.NewNamespace("admin",
       beego.NSRouter("/main", &controllers.MainController{}, "get:Main"),
       beego.NSRouter("/user", &controllers.MainController{}, "get:UserPage"),
       beego.NSRouter("/userUpdate", &controllers.MainController{}, "post:UserUpdate"),
       beego.NSRouter("/post", &controllers.MainController{}, "get:PostPage"),
       beego.NSRouter("/imgUp", &controllers.MainController{}, "post:ImgUp"),
       beego.NSRouter("/resetPwd", &controllers.MainController{}, "post:ResetPwd"),
       )
    beego.AddNamespace(nsAdmin)

    var checkUser = func(ctx *context.Context) {
        /*_, ok := ctx.Input.Session("uid").(int)
        if !ok && ctx.Request.RequestURI != "/login" {
            ctx.Redirect(302, "/login") }*/

    }
    beego.InsertFilter("/*",beego.BeforeRouter,checkUser)
}
