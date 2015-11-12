package controllers

import (
	"github.com/benreic/dingus/persistence"
	"github.com/benreic/dingus/sessions"
	"github.com/benreic/dingus/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func Login(c *gin.Context) {

	submittedPassword := c.PostForm("passwd")
	handle := c.PostForm("username")

	db := persistence.NewDb()

	var hashedPassword string
	var driverId int

	err := db.QueryRow("select handle, driver_id, password from driver where handle = ?", handle).Scan(&handle, &driverId, &hashedPassword)

	if err != nil {
		utils.LogError(err)
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(submittedPassword))
	if err != nil {
		utils.LogError(err)
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}

	session := sessions.Get(c)
	session.Handle = handle
	session.DriverId = driverId
	session.Save(c)

	c.Redirect(http.StatusMovedPermanently, "/dashboard")
}

func SignUp(c *gin.Context) {

	c.HTML(http.StatusOK, "sign-up.tmpl", gin.H{})
}

func ProcessSignUp(c *gin.Context) {

	pwd := c.PostForm("passwd")
	password := []byte(pwd)
	email := c.PostForm("email")
	handle := c.PostForm("username")

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		utils.LogError(err)
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}

	db := persistence.NewDb()
	insertSql := "insert into driver (email, handle, password, created_on) values (?, ?, ?, ?)"
	db.Exec(insertSql, email, handle, hashedPassword, time.Now())

	c.Redirect(http.StatusMovedPermanently, "/")
}
