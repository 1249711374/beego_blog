package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/*
dbhost = 127.0.0.1
dbport = 3306
dbuser = root
dbpassword = 123456
dbname = beego_blog
*/
func init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")

	//"root:111111@tcp(127.0.0.1:3306)/HelloBeego?charset=utf8"
	dburl := dbuser + ":" + dbpassword + "@tcp(" +
		dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	fmt.Println("dburl = ", dburl)
	// set default database
	orm.RegisterDataBase("default", "mysql", dburl, 30)

	// register model
	orm.RegisterModel(new(Link), new(Mood), new(Post), new(Tag), new(TagPost), new(User))
}

// 查询最新的4篇文章
func GetLastestBlog() []*Post {
	post := Post{}
	query := orm.NewOrm().QueryTable(&post).Filter("status", 0)
	count, _ := query.Count()
	var result []*Post
	if count > 0 {
		query.OrderBy("-posttime").Limit(4).All(&result)
	}
	return result
}

// 最多浏览量
func GetHotBlog() []*Post {
	post := Post{}
	query := orm.NewOrm().QueryTable(&post).Filter("status", 0)
	count, _ := query.Count()
	var result []*Post
	if count > 0 {
		query.OrderBy("-views").Limit(4).All(&result)
	}
	return result
}

// 友情连接- 因为不是很多 所以直接ALL()
func GetLinks() []*Link {
	link := Link{}
	query := orm.NewOrm().QueryTable(&link)
	count, _ := query.Count()
	var result []*Link
	if count > 0 {
		query.OrderBy("-rank").All(&result)
	}
	return result
}

//md5哈希
func Md5(buf []byte) string {
	//创建哈希对象
	mymd5 := md5.New()
	//将待加密数据写入哈希对象
	mymd5.Write(buf)
	//获取哈希值
	result := mymd5.Sum(nil)
	return fmt.Sprintf("%x", result)
}
