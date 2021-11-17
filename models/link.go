package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
)

//友情链接
type Link struct {
	Id       int
	Sitename string `orm:"size(80)"`  //网站名称
	Url      string `orm:"size(200)"` //网址
	Rank     int    //排序
}

func (link *Link) TableName() string {
	// 在配置文件中读取(表中 tb_link 和结构体 link 不匹配)
	prefix := beego.AppConfig.String("dbprefix")
	return prefix + "link"
}

func (link *Link) Insert() (rows int, err error) {
	ormer := orm.NewOrm()
	insert, err := ormer.Insert(link)
	return int(insert), err
}

func (link *Link) Delete() error {
	_, err := orm.NewOrm().Delete(link)
	if err != nil {
		log.Println("删除失败")
		return err
	}
	return nil
}

// 读取
func (link *Link) Read(fields ...string) error {
	if err := orm.NewOrm().Read(link, fields...); err != nil {
		return err
	}
	return nil
}

// 更新
func (link *Link) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(link, fields...); err != nil {
		return err
	}
	return nil
}
