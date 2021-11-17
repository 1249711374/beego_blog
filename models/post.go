package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
	"strings"
	"time"
)

//文章
type Post struct {
	Id int
	//用户id
	Userid int
	//作者
	Author string `orm:"size(15)"`
	//标题
	Title string `orm:"size(100)"`
	//标题颜色
	Color string `orm:"size(7)"`
	//文章内容
	Content string `orm:"type(text)"`
	//标签名称
	Tags string `orm:"size(100)"`
	//浏览量
	Views int
	//状态
	Status int
	//发表时间
	Posttime time.Time `orm:"type(datetime)"`
	//更新时间
	Updated time.Time `orm:"type(datetime)"`
	//是否置顶
	Istop int
	//封面
	Cover string `orm:"size(70)"`
}

func (link *Post) TableName() string {
	// 在配置文件中读取(表中 tb_link 和结构体 link 不匹配)
	prefix := beego.AppConfig.String("dbprefix")
	return prefix + "post"
}

func (link *Post) Insert() (rows int, err error) {
	ormer := orm.NewOrm()
	insert, err := ormer.Insert(link)
	return int(insert), err
}

func (link *Post) Delete() error {
	_, err := orm.NewOrm().Delete(link)
	if err != nil {
		log.Println("删除失败")
		return err
	}
	return nil
}

// 读取
func (link *Post) Read(fields ...string) error {
	if err := orm.NewOrm().Read(link, fields...); err != nil {
		return err
	}
	return nil
}

// 更新
func (link *Post) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(link, fields...); err != nil {
		return err
	}
	return nil
}

func (post *Post) TagsLink() string {
	if post.Tags == "" {
		return ""
	}
	trim := strings.Trim(post.Tags, ",")
	return trim
}

func (post *Post) Link() string {
	// /index:page:int.html 首页路径
	// /article/2
	return "/article/" + strconv.Itoa(post.Id)
}

func (post *Post) ColorTitle() string {
	if post.Color != "" {
		return fmt.Sprintf("<span style='color:%s'>%s</span>", post.Color, post.Title)
	}
	return post.Title
}

func (post *Post) Excerpt() string {
	return post.Content
}

// 获取当前文章 的上一篇文章和下一篇文章
func (post *Post) GetPreAndNext() (pre, next *Post) {
	pre = &Post{}
	orm.NewOrm().QueryTable(new(Post)).OrderBy("-id").Filter("id__lt", post.Id).Filter("status", 0).Limit(1).One(pre)
	next = &Post{}
	orm.NewOrm().QueryTable(new(Post)).OrderBy("id").Filter("id__gt", post.Id).Filter("status", 0).Limit(1).One(next)
	return pre, next
}
