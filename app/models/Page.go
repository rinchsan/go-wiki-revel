package models

type Page struct {
	ID    int64  `gorm:"primary_key"`
	Title string `sql:"not null"`
	Body  []byte
}

func (p *Page) create() {
	if DB.NewRecord(p) {
		DB.Create(p)
	}
}

func (p *Page) Update(body string) {
	p.Body = []byte(body)
	DB.Save(p)
}

func LoadOrCreatePage(title string) (*Page, bool) {
	var pages []*Page
	DB.Where("title = ?", title).First(&pages)
	if len(pages) > 0 {
		return pages[0], false
	}
	p := &Page{Title: title}
	p.create()
	return p, true
}

func GetAllPages() (pages []*Page) {
	DB.Find(&pages)
	return
}
