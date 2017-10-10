package controllers

import (
	"go-wiki-revel/app/models"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Root() revel.Result {
	pages, err := models.GetAllPages()
	if err != nil {
		return c.RenderError(err)
	}
	return c.Render(pages)
}
