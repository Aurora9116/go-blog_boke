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

	// AddPost GetAllPost getPost 博客操作
	AddPost(post *model.Post)
	GetAllPost() []model.Post
	GetPost(pid int) model.Post
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
	_ = db.AutoMigrate(&model.User{}, &model.Post{})
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

func (mgr *manager) AddPost(post *model.Post) {
	mgr.db.Create(post)
}

func (mgr *manager) GetAllPost() []model.Post {
	var posts = make([]model.Post, 10)
	mgr.db.Find(&posts)
	return posts
}

func (mgr *manager) GetPost(pid int) model.Post {
	var post model.Post
	mgr.db.Find(&post, pid)
	return post
}
