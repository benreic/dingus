package controllers

import (
	"github.com/benreic/dingus/sessions"
	"github.com/benreic/dingus/utils"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func Hello(c *gin.Context) {

	session := sessions.Get(c)
	if session.IsLoggedIn() {
		c.Redirect(http.StatusMovedPermanently, "/dashboard")
		return
	}

	content := utils.Render("templates/index.tmpl", gin.H{"handle": session.Handle, "loggedIn": false})
	c.HTML(http.StatusOK, "master.tmpl", gin.H{"content": template.HTML(content)})
}
