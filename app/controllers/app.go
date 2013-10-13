package controllers

import (
	"github.com/dougnukem/revel-helloworld/app/models"
	"github.com/robfig/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	user := models.User{Id: 1, Name: "Hal"}
	return c.Render(user)
}
