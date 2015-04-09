package models

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
    "time"
)

func init() {
    dbhost := beego.AppConfig.String("dbhost")
    dbport := beego.AppConfig.String("dbport")
    dbuser := beego.AppConfig.String("dbuser")
    dbpassword := beego.AppConfig.String("dbpassword")
    dbname := beego.AppConfig.String("dbname")
    if dbport == "" {
        dbport = "3306"
    }
    dburl := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
    orm.RegisterModel(new(Post), new(User), new(Message))
    orm.RegisterDataBase("default", "mysql", dburl)
    if beego.AppConfig.String("runmode") == "dev" {
        orm.Debug = true
    }
}

type User struct {
    Id int64 `json:"id"`
    UserName string `orm:"size(50)" json:"userName"`
    SiteTitle string `orm:"size(50)" json:"siteTitle"`
    Password string `orm:"size(100)" json:"-"`
    Thumb    string `orm:"size(500)" json:"thumb"`
    HeadBgPic string `json:"headBgPic"`
    AboutMe string `orm:"size(2000)" json:"aboutMe"`
}

func (this *User) Query() orm.QuerySeter {
    return orm.NewOrm().QueryTable(this)
}

func (this *User) Update(){
    o := orm.NewOrm()
    o.Update(this)
}

type Post struct {
    Id int64 `json:"id"`
    Title string `json:"title"`
    Content string `orm:"type(text)" json:"content"`
    Thumb string `json:"thumb"`
    Tag string `json:"tag"`
    PublishAt time.Time `orm:"auto_now_add;type(datetime)" json:"publishAt"`
}

func (this *Post) Read() error {
    o := orm.NewOrm()
    return o.Read(this)
}

func (this *Post) Insert() error {
    if _, err := orm.NewOrm().Insert(this); err != nil {
        return err
    }
    return nil
}

func (this *Post) Update(){
    o := orm.NewOrm()
    o.Update(this)
}

func (this *Post) Delete() error{
    if _, err := orm.NewOrm().Delete(this); err != nil {
        return err
    }
    return nil
}

func (this *Post) Query() orm.QuerySeter {
    return orm.NewOrm().QueryTable(this)
}

type Message struct {
    Id int64 `json:"id"`
    GuestName string `json:"guestName"`
    Content string `orm:"type(text)" json:"content"`
    PostId int64 `json:"postId"`
    PostTitle string `json:"postTitle"`
    CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"createdAt"`
}

func (this *Message) Insert() error {
    if _, err := orm.NewOrm().Insert(this); err != nil {
        return err
    }
    return nil
}

func (this *Message) Delete() error{
    if _, err := orm.NewOrm().Delete(this); err != nil {
        return err
    }
    return nil
}

func (this *Message) Query() orm.QuerySeter {
    return orm.NewOrm().QueryTable(this)
}