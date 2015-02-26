package controllers

import (
	"github.com/astaxie/beego"
	"github.com/igordonshaw/inote/models"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type ResEntity struct {
	Success bool `json:"success"`
	Msg string `json:"msg"`
	Data interface {} `json:"data"`
}

type BaseController struct {
	beego.Controller
}

type MainController struct {
	BaseController
}

func (this *MainController) LoginPage(){
	this.TplNames = "login.html"
}

func (this *MainController) Login(){
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"

	if beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd{
		maxAge := 0
		if autoLogin {
			maxAge = 1 << 31 -1
		}

		this.Ctx.SetCookie("uname", uname, maxAge, "/")
	}

	this.Redirect("/", 301)
	return
}

func (this *MainController) Get() {
	o := orm.NewOrm()
	var posts []orm.Params
	_, err := o.Raw("SELECT id,title,publish_at FROM "+beego.AppConfig.String("dbprefix")+"post order by publish_at desc").Values(&posts)
	if err != nil {
		beego.Trace("no post")
	}
	this.Data["Posts"] = posts

	qsCategories := new(models.Category)
	var categories []*models.Category
	qsCategories.Query().All(&categories)
	this.Data["Categories"] = categories
	this.TplNames = "index.html"
}

type CategoryController struct {
	BaseController
}


type PostController struct {
	BaseController
}

func (this *PostController) One(){
	id,err :=strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err != nil {
		beego.Error(err)
	}
	qsPost := new(models.Post)
	post := models.Post{Id: int64(id)}
	qsPost.Query().RelatedSel().One(&post)
	this.Data["Post"] = post
	this.TplNames = "postdetail.html"
}