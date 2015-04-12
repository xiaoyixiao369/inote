package controllers

import (
	"github.com/astaxie/beego"
	"github.com/igordonshaw/inote/models"
	"strconv"
    "strings"
    "time"
    "encoding/json"
    "crypto/md5"
    "encoding/hex"
    "github.com/astaxie/beego/orm"
)

var IMG_EXT = []string{"jpg","jpeg","png","bmp","webp","JPG","JPEG","PNG","BMP","WEBP"}
var PAGE_SIZE int
var LOGIN_LOCK_TIMES int
var RELOGIN_PERIOD int64
var LOGIN_COUNT = 0
var LAST_LOGIN_TIME = time.Now().Unix()

func init(){
    pageSize, err := beego.AppConfig.Int("pagesize")
    if err != nil {
        PAGE_SIZE = 10
    }
    PAGE_SIZE = pageSize

    loginLockTimes, err := beego.AppConfig.Int("loginlocktimes")
    if err != nil {
        LOGIN_LOCK_TIMES = 5
    }
    LOGIN_LOCK_TIMES = loginLockTimes

    reLoginPeriod, err := beego.AppConfig.Int64("reloginperiod")
    if err != nil {
        reLoginPeriod = 30
    }
    RELOGIN_PERIOD = reLoginPeriod
}

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

func (this *MainController) MessagePage(){
    this.TplNames = "admin/message.html"
}

func (this *MainController) WritePage(){
    this.TplNames = "admin/write.html"
}

func (this *UserControlelr) Author(){
    qsUser := new(models.User)
    user := models.User{Id: int64(1)}
    qsUser.Query().Filter("id", 1).One(&user)
    this.Data["json"] = user
    this.ServeJson()
}

func (this *MainController) ValidUser(){
    res := &ResEntity{}
    if (time.Now().Unix() - LAST_LOGIN_TIME < RELOGIN_PERIOD){
        LOGIN_COUNT++
    } else {
        LOGIN_COUNT = 0
    }

    if LOGIN_COUNT >= LOGIN_LOCK_TIMES {
        res.Success = false
        if LOGIN_COUNT > LOGIN_LOCK_TIMES + 3{
            res.Msg = "悟空，不要调皮了"
        } else {
            res.Msg = "连续尝试"+strconv.Itoa(LOGIN_COUNT)+"次登录失败, 请隔"+strconv.Itoa(int(RELOGIN_PERIOD))+"秒后再登录"
        }
        this.Data["json"] = res
        this.ServeJson()
        return;
    }

    h := md5.New()
    h.Write([]byte(string(this.Ctx.Input.RequestBody)))
    inputPwd := hex.EncodeToString(h.Sum(nil))
    qsUser := new(models.User)
    userDb := models.User{Id: 1}
    qsUser.Query().Filter("id", int64(1)).One(&userDb)
    if userDb.Password != inputPwd {
        LAST_LOGIN_TIME = time.Now().Unix()
        res.Success = false
        res.Msg = "登录失败"
        this.Data["json"] = res
        this.ServeJson()
        return;
    }
    /*maxAge := 0
    maxAge = 1 << 31 -1
    this.Ctx.SetCookie("username", "inote", maxAge, "/")*/
    LOGIN_COUNT = 0
    LAST_LOGIN_TIME = time.Now().Unix()
    this.SetSession("inote", 1)
	res.Success = true
    res.Msg = "登录成功"
    this.Data["json"] = res
    this.ServeJson()
	return
}

func (this *MainController) Logout(){
    this.DelSession("inote")
    res := &ResEntity{true, "退出成功", nil}
    this.Data["json"] = res
    this.ServeJson()
    return
}


func (this *MainController) Get() {
    this.TplNames = "main.html"
}

func (this *MainController) UserUpdate(){
    var userFront models.User
    err := json.Unmarshal(this.Ctx.Input.RequestBody, &userFront)
    if err != nil {
        beego.Error("invalid user," + err.Error())
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
    page,err :=strconv.Atoi(this.Ctx.Input.Param(":page"))
    if err != nil {
        beego.Error(err)
    }
    var posts []models.Post
    qb, _ := orm.NewQueryBuilder("mysql")

    qb.Select("id","title","tag","publish_at").
    From("post").
    OrderBy("publish_at").Desc().
    Limit(PAGE_SIZE).Offset((page - 1) * PAGE_SIZE)
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

type FrontPost struct {
    Id string `json:"id"`
    Title string `json:"title"`
    Tag string `json:"tag"`
    Content string `json:"content"`
}

func (this *MainController) SavePost(){
    res := &ResEntity{}
    var frontPost FrontPost
    err := json.Unmarshal(this.Ctx.Input.RequestBody, &frontPost)
    if err != nil {
        beego.Error("invalid post," + err.Error())
        res.Success = false
        res.Msg = "无效的内容"
        this.Data["json"] = res
        this.ServeJson()
        return
    }

    post := &models.Post{
        Title: frontPost.Title,
        Tag: frontPost.Tag,
        Content:frontPost.Content,
    }

    if "" == post.Title {
        post.Title = "未命名"
    }

    if "" == post.Tag {
        post.Tag = "默认标签"
    }

    if "" == frontPost.Id {
        post.Insert()
    } else {
        postId, err := strconv.Atoi(frontPost.Id)
        if err != nil {
            beego.Error("invalid post id:", err.Error())
        }
        post.Id = int64(postId)
        postDb := &models.Post{Id: post.Id}
        postDb.Read()
        post.PublishAt = time.Now()
        post.Update()
    }

    res.Success = true
    res.Msg = "保存成功"
    this.Data["json"] = res
    this.ServeJson()
    return
}

func (this *MainController) DeletePost(){
    res := &ResEntity{}
    id,err :=strconv.Atoi(this.Ctx.Input.Param(":id"))
    if err != nil {
        beego.Error(err)
    }
    post := models.Post{Id: int64(id)}
    err = post.Delete()
    if err != nil {
        beego.Error("delete post error:", err.Error())
        res.Success = false
        res.Msg = "删除失败"
        this.Data["json"] = res
        this.ServeJson()
        return
    }

    o := orm.NewOrm()
    o.Raw("DELETE FROM message WHERE post_id=?", id).Exec()

    res.Success = true
    res.Msg = "删除成功"
    this.Data["json"] = res
    this.ServeJson()
    return
}

type FrontMessage struct {
    PostId string `json:"postId"`
    PostTitle string `json:"postTitle"`
    GuestName string `json:"guestName"`
    Content string `json:"content"`
}

func(this *PostController) SubmitMsg(){
    res := &ResEntity{}
    var frontMessage FrontMessage
    err := json.Unmarshal(this.Ctx.Input.RequestBody, &frontMessage)
    if err != nil {
        beego.Error("invalid message," + err.Error())
        res.Success = false
        res.Msg = "无效的留言"
        this.Data["json"] = res
        this.ServeJson()
        return
    }

    if "" == frontMessage.GuestName {
        frontMessage.GuestName = "佚名"
    }
    postId,err :=strconv.Atoi(frontMessage.PostId)
    if err != nil {
        beego.Error(err)
    }
    message := &models.Message{
        GuestName: frontMessage.GuestName,
        Content: frontMessage.Content,
        PostId: int64(postId),
        PostTitle: frontMessage.PostTitle,
    }

    err = message.Insert()
    if err != nil {
        res.Success = false
        res.Msg = "添加留言失败"
        this.Data["json"] = res
        this.ServeJson()
        return
    }
    res.Success = true
    res.Msg = "添加留言成功"
    this.Data["json"] = res
    this.ServeJson()
    return

}

type ResPost struct {
    Post *models.Post `json:"post"`
    Messages *[]models.Message `json:"messages"`
}

func (this *PostController) OnePost(){
    res := &ResEntity{}
	id,err :=strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err != nil {
		beego.Error(err)
	}
	qsPost := new(models.Post)
	post := models.Post{Id: int64(id)}
    if id == 0 {
        // last post
        qsPost.Query().OrderBy("-PublishAt").Limit(1).One(&post)
    } else {
        qsPost.Query().RelatedSel().Filter("id", id).One(&post)
    }

    /*if post == nil {
        res.Success = false
        res.Msg = "还没有内容"
        this.Data["json"] = res
        this.ServeJson()
        return
    }*/
    messages := []models.Message{}
    qsMessages := new(models.Message)
    qsMessages.Query().Filter("PostId", post.Id).OrderBy("-CreatedAt").All(&messages)
    resPost := &ResPost{
        Post: &post,
        Messages: &messages,
    }
    res.Success = true
    res.Data = resPost
	this.Data["json"] = res
    this.ServeJson()
    return
}

func (this *MainController) ListMessage(){
    page,err :=strconv.Atoi(this.Ctx.Input.Param(":page"))
    if err != nil {
        beego.Error(err)
    }
    var messages []models.Message
    qb, _ := orm.NewQueryBuilder("mysql")

    qb.Select("id","guest_name","content","post_title", "created_at").
    From("message").
    OrderBy("created_at").Desc().
    Limit(PAGE_SIZE + 10).Offset((page - 1) * PAGE_SIZE)
    sql := qb.String()

    o := orm.NewOrm()
    o.Raw(sql).QueryRows(&messages)

    this.Data["json"] = messages
    this.ServeJson()
    return
}

func (this *MainController) DeleteMessage(){
    res := &ResEntity{}
    id,err :=strconv.Atoi(this.Ctx.Input.Param(":id"))
    if err != nil {
        beego.Error(err)
    }
    message := models.Message{Id: int64(id)}
    err = message.Delete()
    if err != nil {
        beego.Error("delete post error:", err.Error())
        res.Success = false
        res.Msg = "删除失败"
        this.Data["json"] = res
        this.ServeJson()
        return
    }

    res.Success = true
    res.Msg = "删除成功"
    this.Data["json"] = res
    this.ServeJson()
    return
}

func (this *MainController) ImgUp() {
    _, fileHeder, err := this.GetFile("editormd-image-file")
    if err != nil {
        beego.Error(err.Error())
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
    err = this.SaveToFile("editormd-image-file", beego.AppPath + "/" + beego.AppConfig.String("uploaddir") + fileNewName)
    if err != nil {
        beego.Error(err.Error())
    }

    res := &ResEntity{true, "", "/" + beego.AppConfig.String("uploaddir") + fileNewName}
    this.Data["json"] = res
    this.ServeJson()
}


type ErrorController struct {
    beego.Controller
}

func (c *ErrorController) Error404() {
}

func (c *ErrorController) Error403() {
}
