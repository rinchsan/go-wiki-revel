package models

import (
	"io/ioutil"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile("go-wiki-revel/public/data/"+filename, p.Body, 0600)
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
