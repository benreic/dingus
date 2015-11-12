package controllers

import (
	"github.com/benreic/dingus/config"
	"github.com/benreic/dingus/sessions"
	"github.com/benreic/dingus/utils"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func Dashboard(c *gin.Context) {
	session := sessions.Get(c)
	content := utils.Render("templates/dashboard.tmpl", gin.H{"handle": session.Handle, "client_id": config.ClientId()})
	c.HTML(http.StatusOK, "master.tmpl", gin.H{"content": template.HTML(content)})
}

func Logout(c *gin.Context) {

	session := sessions.Get(c)
	session.DriverId = 0
	session.Handle = ""
	session.Save(c)
	c.Redirect(http.StatusMovedPermanently, "/")
}
