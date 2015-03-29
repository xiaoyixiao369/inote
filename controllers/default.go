package controllers

import (
	"github.com/astaxie/beego"
	"github.com/igordonshaw/inote/models"
	"strconv"
    "fmt"
    "strings"
    "time"
    "encoding/json"
)

var IMG_EXT = []string{"jpg","jpeg","png","JPG","JPEG","PNG"}

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

type UserControlelr struct {
    BaseController
}

type CategoryController struct {
    BaseController
}


type PostController struct {
    BaseController
}

func (this *MainController) LoginPage(){
	this.TplNames = "admin/login.html"
}

func (this *MainController) Main(){
    this.TplNames = "admin/main.html"
}

func (this *MainController) UserPage(){
    this.TplNames = "admin/user.html"
}

func (this *MainController) CategoryPage(){
    this.TplNames = "admin/category.html"
}

func (this *MainController) PostPage(){
    this.TplNames = "admin/post.html"
}

func (this *UserControlelr) Author(){
    qsUser := new(models.User)
    user := models.User{Id: int64(1)}
    qsUser.Query().Filter("id", 1).One(&user)
    this.Data["json"] = user
    this.ServeJson()
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

func (this *MainController) UserUpdate(){
    var userFront models.User
    err := json.Unmarshal(this.Ctx.Input.RequestBody, &userFront)
    if err != nil {
        fmt.Println("invalid user," + err.Error())
    }

    qsUser := new(models.User)
    userDb := models.User{Id: int64(userFront.Id)}
    qsUser.Query().Filter("id", int64(userFront.Id)).One(&userDb)

    userFront.Password = userDb.Password
    userFront.Update()
    res := &ResEntity{true, "修改成功", nil}
    this.Data["json"] = res
    this.ServeJson()
    return
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

func (this *MainController) ImgUp() {

    _, fileHeder, err := this.GetFile("avatar")
    if err != nil {
        fmt.Println(err.Error())
    }
    fileName := fileHeder.Filename

    if strings.Index(fileName, ".") <= 0 {
        res := &ResEntity{false, "错误的图片文件!", ""}
        this.Data["json"] = res
        this.ServeJson()
        return
    }

    strs := strings.Split(fileName, ".")
    ext := strs[len(strs) - 1]

    isExtPass := false

    for _, allowedExt := range IMG_EXT {
        if allowedExt == ext {
            isExtPass = true
            break;
        }
    }

    if !isExtPass {
        res := &ResEntity{false, "不支持的图片格式!", ""}
        this.Data["json"] = res
        this.ServeJson()
        return
    }

    fileNewName := strconv.FormatInt(time.Now().Unix(), 10) + "." + ext
    err = this.SaveToFile("avatar", beego.AppPath + "/" + beego.AppConfig.String("uploaddir") + fileNewName)
    if err != nil {
        fmt.Println(err.Error())
    }

    res := &ResEntity{true, "", "/" + beego.AppConfig.String("uploaddir") + fileNewName}
    this.Data["json"] = res
    this.ServeJson()
}
