package models

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

func LoadPage(title string) (p *Page) {
	var pages []*Page
	DB.Where("title = ?", title).First(&pages)
	if len(pages) > 0 {
		return pages[0]
	} else {
		return nil
	}
}

func GetAllPages() (pages []*Page) {
	DB.Find(&pages)
	return
}
