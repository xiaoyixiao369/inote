package controllers

import (
	"github.com/astaxie/beego"
	"github.com/igordonshaw/inote/models"
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

type CategoryController struct {
    BaseController
}


type PostController struct {
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
    posts := []models.Post{}
    qsPost := new(models.Post)
    qsPost.Query().OrderBy("-PublishAt").All(&posts)
    if len(posts) == 0 {
        flash := beego.NewFlash()
        flash.Notice("还没有文章哦!")
        flash.Store(&this.Controller)
    }
    this.Data["Posts"] = posts

    if categories, err := Categories(); err == nil {
        this.Data["Categories"] = categories
    }

    this.Layout = "index.html"
    this.TplNames = "posts.tpl"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["Header"] = "header.tpl"
    this.LayoutSections["Sidebar"] = "sidebar.tpl"
}

func Categories() ([]*models.Category, error) {
    qsCategories := new(models.Category)
    var categories []*models.Category
    if _, err := qsCategories.Query().All(&categories); err != nil {
        return nil ,err;
    }
    return categories, nil
}

func (this *PostController) One(){
	id,err :=strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err != nil {
		beego.Error(err)
	}
	qsPost := new(models.Post)
	post := models.Post{Id: int64(id)}
	qsPost.Query().RelatedSel().Filter("id", id).One(&post)
	this.Data["Post"] = post
    this.Layout = "postdetail.html"
    this.TplNames = "post.tpl"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["Header"] = "header.tpl"
}

func (this *PostController) Category(){
    id,err :=strconv.Atoi(this.Ctx.Input.Param(":id"))
    if err != nil {
        beego.Error(err)
    }
    posts := []models.Post{}
    qsPost := new(models.Post)
    qsPost.Query().Filter("Category__id", id).OrderBy("-PublishAt").RelatedSel().All(&posts)
    if len(posts) == 0 {
        flash := beego.NewFlash()
        flash.Notice("该分类下还没有文章哦!")
        flash.Store(&this.Controller)
    }
    this.Data["Posts"] = posts

    if categories, err := Categories(); err == nil {
        this.Data["Categories"] = categories
    }
    this.Layout = "index.html"
    this.TplNames = "posts.tpl"
    this.LayoutSections = make(map[string]string)
    this.LayoutSections["Header"] = "header.tpl"
    this.LayoutSections["Sidebar"] = "sidebar.tpl"
}