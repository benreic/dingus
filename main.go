package main

import (
	"github.com/benreic/dingus/config"
	"github.com/benreic/dingus/persistence"
	"github.com/benreic/dingus/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {

	logger := utils.NewLogger()
	logger.Println("Started")

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", Hello)
	router.GET("/automatic-redirect", AutomaticRedirect)
	router.GET("/sign-up", SignUp)
	router.POST("/sign-up", ProcessSignUp)
	router.Static("/public", "./public")

	// Listen and server on 0.0.0.0:8080
	err := router.Run(":" + config.Port())

	logger.Println("Stopping")

	if err != nil {
		logger.Println(err.Error())
	}
}

func Hello(c *gin.Context) {

	content := utils.Render("templates/index.tmpl", gin.H{"name": "Dingus"})
	c.HTML(http.StatusOK, "master.tmpl", gin.H{"content": template.HTML(content)})
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
		panic(err)
	}

	db := persistence.NewDb()
	insertSql := "insert into driver (email, handle, password, created_on) values (?, ?, ?, ?)"
	db.Exec(insertSql, email, handle, hashedPassword, time.Now())

	c.HTML(http.StatusOK, "sign-up.tmpl", gin.H{})
}

func AutomaticRedirect(c *gin.Context) {

	oauthUrl := "https://accounts.automatic.com/oauth/access_token"
	postData := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {c.Query("code")},
		"client_id":     {config.ClientId()},
		"client_secret": {config.ClientSecret()},
	}

	resp, _ := http.PostForm(oauthUrl, postData)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	c.HTML(http.StatusOK, "master.tmpl", gin.H{"content": string(body)})
}
