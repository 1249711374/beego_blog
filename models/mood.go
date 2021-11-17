package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

//碎言碎语
type Mood struct {
	Id int
	//说说内容
	Content string `orm:"type(text)"`
	//封面路径
	Cover string `orm:"size(70)"`
	//发表时间
	Posttime time.Time `orm:type(datetime)`
}

func (link *Mood) TableName() string {
	// 在配置文件中读取(表中 tb_link 和结构体 link 不匹配)
	prefix := beego.AppConfig.String("dbprefix")
	return prefix + "mood"
}

func (link *Mood) Insert() (rows int, err error) {
	ormer := orm.NewOrm()
	insert, err := ormer.Insert(link)
	return int(insert), err
}

func (link *Mood) Delete() error {
	_, err := orm.NewOrm().Delete(link)
	if err != nil {
		log.Println("删除失败")
		return err
	}
	return nil
}

// 读取
func (link *Mood) Read(fields ...string) error {
	if err := orm.NewOrm().Read(link, fields...); err != nil {
		return err
	}
	return nil
}

// 更新
func (link *Mood) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(link, fields...); err != nil {
		return err
	}
	return nil
}
