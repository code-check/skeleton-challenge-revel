package controllers

import "github.com/revel/revel"

type HelloWorld struct {
  World string
}

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	var h HelloWorld
	h.World = "Hello"
	return c.RenderJson(h)
}