package models

import (
	"io/ioutil"
)

type Page struct {
	ID    int64  `gorm:"primary_key"`
	Title string `sql:"not null"`
	Body  []byte
}

func (p *Page) SaveOrUpdate() {
	if DB.NewRecord(p) {
		DB.Create(p)
	} else {
		DB.Save(p)
	}
}

func LoadPage(title string) (*Page, error) {
	filename := "go-wiki-revel/public/data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func GetAllPages() (pages []*Page) {
	DB.Find(&pages)
	return
}
