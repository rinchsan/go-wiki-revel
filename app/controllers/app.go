package controllers

import (
	"go-wiki-revel/app/models"
	"regexp"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Root() revel.Result {
	pages := models.GetAllPages()
	c.ViewArgs["pages"] = pages
	return c.Render()
}

func (c App) executeAction(action func(string) revel.Result, title string) revel.Result {
	rep := regexp.MustCompile("^[a-zA-Z0-9]+$")
	matched := rep.Match([]byte(title))
	if matched {
		return action(title)
	}
	return c.NotFound("Invalid title")
}

func (c App) view(title string) revel.Result {
	p, created := models.LoadOrCreatePage(title)
	if created {
		return c.Redirect("/edit/" + p.Title)
	}
	c.ViewArgs["page"] = p
	return c.Render()
}

func (c App) edit(title string) revel.Result {
	p, _ := models.LoadOrCreatePage(title)
	c.ViewArgs["page"] = p
	return c.Render()
}

func (c App) save(title string) revel.Result {
	body := c.Request.FormValue("body")
	p, _ := models.LoadOrCreatePage(title)
	p.Update(body)
	return c.Redirect("/view/" + title)
}

func (c App) View() revel.Result {
	title := c.Params.Route.Get("title")
	return c.executeAction(c.view, title)
}

func (c App) Edit() revel.Result {
	title := c.Params.Route.Get("title")
	return c.executeAction(c.edit, title)
}

func (c App) Save() revel.Result {
	title := c.Params.Route.Get("title")
	return c.executeAction(c.save, title)
}

func (c App) NewGet() revel.Result {
	return c.RenderTemplate("App/new.html")
}

func (c App) NewPost() revel.Result {
	title := c.Request.FormValue("title")
	return c.Redirect("/view/" + title)
}
