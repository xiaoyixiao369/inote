package controllers

import (
	"github.com/astaxie/beego"
	"github.com/igordonshaw/inote/models"
	"strconv"
    "fmt"
    "strings"
    "time"
    "encoding/json"
    "crypto/md5"
    "encoding/hex"
    "github.com/astaxie/beego/orm"
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
    this.TplNames = "main.html"
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

func (this *MainController) ResetPwd(){
    h := md5.New()
    h.Write([]byte(string(this.Ctx.Input.RequestBody)))
    newPwd := hex.EncodeToString(h.Sum(nil))
    qsUser := new(models.User)
    userDb := models.User{Id: 1}
    qsUser.Query().Filter("id", int64(1)).One(&userDb)
    userDb.Password = newPwd
    userDb.Update()

    res := &ResEntity{true, "修改成功", nil}
    this.Data["json"] = res
    this.ServeJson()
    return;
}

func (this *PostController) ListPosts(){
    var posts []models.Post
    qb, _ := orm.NewQueryBuilder("mysql")

    qb.Select("id","title","tag","publish_at").
    From("post").
    OrderBy("publish_at").Desc().
    Limit(10).Offset(0)
    sql := qb.String()

    o := orm.NewOrm()
    o.Raw(sql).QueryRows(&posts)

    this.Data["json"] = posts
    this.ServeJson()
    return
}

func (this *PostController) Posts(){
    posts := []models.Post{}
    qsPost := new(models.Post)
    qsPost.Query().OrderBy("-PublishAt").All(&posts)
    this.Data["json"] = posts
    this.ServeJson()
    return
}

func (this *PostController) OnePost(){
	id,err :=strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err != nil {
		beego.Error(err)
	}
	qsPost := new(models.Post)
	post := models.Post{Id: int64(id)}
	qsPost.Query().RelatedSel().Filter("id", id).One(&post)
	this.Data["json"] = post
    this.ServeJson()
    return
}

func (this *MainController) ImgUp() {
    _, fileHeder, err := this.GetFile("avatar")
    if err != nil {
        fmt.Println(err.Error())
        res := &ResEntity{false, "服务器错误",nil}
        this.Data["json"] = res
        this.ServeJson()
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
