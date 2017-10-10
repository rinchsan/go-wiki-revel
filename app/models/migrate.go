package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

var DB *gorm.DB

func InitDB() {
	user := revel.Config.StringDefault("db.user", "root")
	name := revel.Config.StringDefault("db.name", "go_wiki_revel")
	dsn := user + "@/" + name + "?charset=utf8&parseTime=True"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	migrate()
}

func migrate() {
	DB.AutoMigrate(&Page{})
}
