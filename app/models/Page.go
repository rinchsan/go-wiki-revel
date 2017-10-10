package models

import (
	"io/ioutil"
	"regexp"
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

func GetAllPages() ([]*Page, error) {
	fileInfos, err := ioutil.ReadDir("go-wiki-revel/public/data")
	if err != nil {
		return nil, err
	}
	var pages []*Page
	for _, fileInfo := range fileInfos {
		title := fileInfo.Name()
		rep, err := regexp.Compile(".txt$")
		if err != nil {
			continue
		}
		title = rep.ReplaceAllString(title, "")
		page, err := LoadPage(title)
		if err != nil {
			return nil, err
		}
		pages = append(pages, page)
	}
	return pages, nil
}
