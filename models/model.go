package models

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
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
    dburl := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
    orm.RegisterDataBase("default", "mysql", dburl)
    orm.RegisterModel(new(Category), new(Post))
    if beego.AppConfig.String("runmode") == "dev" {
        orm.Debug = true
    }
}

//返回带前缀的表名
func TableName(str string) string {
    return fmt.Sprintf("%s%s", beego.AppConfig.String("dbprefix"), str)
}

type Category struct {
    Id       int64     `json:"id"`
    Name     string    `orm:"size(100)" json:"name"`
}

func (c *Category) TableName() string {
    return TableName("category")
}

func (c *Category) Insert() error {
    if _, err := orm.NewOrm().Insert(c); err != nil {
        return err
    }
    return nil
}

func (c *Category) Read(fields ...string) error {

    if err := orm.NewOrm().Read(c, fields...); err != nil {
        return err
    }
    return nil
}

func (c *Category) Query() orm.QuerySeter {
    return orm.NewOrm().QueryTable(c)
}



type Post struct {
    Id int64 `json:"id"`
    Title string `json:"title"`
    PublishAt string `orm:"auto_now;type(datetime)" json:"publishAt"`
    Content string `orm:"type(text)" json:"content"`
    Thumb string `json:"thumb"`
    Category *Category `orm:"rel(fk)" json:"category"`
}

func (a *Post) TableName() string {
    return TableName("post")
}

func (a *Post) Insert() error {
    if _, err := orm.NewOrm().Insert(a); err != nil {
        return err
    }
    return nil
}

func (a *Post) Read(fields ...string) error {
    if err := orm.NewOrm().Read(a, fields...); err != nil {
        return err
    }
    a.Category.Read()
    return nil
}

func (a *Post) Query() orm.QuerySeter {
    return orm.NewOrm().QueryTable(a)
}