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
        beego.NSRouter("/posts/list/:page", &controllers.PostController{}, "get:ListPosts"),
        beego.NSRouter("/posts/:id", &controllers.PostController{}, "get:OnePost"),
        beego.NSRouter("/submitMsg", &controllers.PostController{}, "post:SubmitMsg"),
    )
    beego.AddNamespace(nsApi)

    nsAdmin :=
       beego.NewNamespace("admin",
       beego.NSRouter("/main", &controllers.MainController{}, "get:Main"),
       beego.NSRouter("/user", &controllers.MainController{}, "get:UserPage"),
       beego.NSRouter("/userUpdate", &controllers.MainController{}, "post:UserUpdate"),
       beego.NSRouter("/post", &controllers.MainController{}, "get:PostPage"),
       beego.NSRouter("/post/save", &controllers.MainController{}, "post:SavePost"),
       beego.NSRouter("/post/delete/:id", &controllers.MainController{}, "delete:DeletePost"),
       beego.NSRouter("/message", &controllers.MainController{}, "get:MessagePage"),
       beego.NSRouter("/message/list/:page", &controllers.MainController{}, "get:ListMessage"),
       beego.NSRouter("/message/delete/:id", &controllers.MainController{}, "delete:DeleteMessage"),
       beego.NSRouter("/write", &controllers.MainController{}, "get:WritePage"),
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
