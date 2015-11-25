package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Me() revel.Result {
	name := "csr"
	return c.Render(name)
}

func (c App) MeNumber(id int) revel.Result {
	return c.Render(id)
}

func (c App) MePost() revel.Result {
	return c.Render()
}

func (c App) MePostSet() revel.Result {
	password := c.Params.Get("password")
	if password == "1234" {
		return c.Redirect(App.Index)
	} else {
		return c.Redirect(App.MePost)
	}
}
