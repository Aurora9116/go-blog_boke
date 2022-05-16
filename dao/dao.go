package dao

import (
	"go-blog_boke/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Manager interface {
	Register(user *model.User)
	Login(username string) model.User
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := "go-blog_boke:123456@tcp(127.0.0.1:3306)/go-blog_boke?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fatal to init db:", err)
	}
	Mgr = &manager{db: db}
	_ = db.AutoMigrate(&model.User{})
	//fmt.Printf("%T", Mgr.(*manager).db)

}

func (mgr *manager) Register(user *model.User) {
	mgr.db.Create(user)
}

func (mgr *manager) Login(username string) model.User {
	var user model.User
	mgr.db.Where("username=?", username).First(&user)
	return user
}
