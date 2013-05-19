package controllers

import (
	"github.com/robfig/revel"
	"helloworld/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	user := models.User{Id: 1, Name: "Hal"}
	return c.Render(user)
}
