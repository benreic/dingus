package main

import (
	"github.com/benreic/dingus/config"
	"github.com/benreic/dingus/persistence"
	"github.com/benreic/dingus/sessions"
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

	utils.Log("Dingus starting")

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.tmpl")
	router.Static("/public", "./public")

	public := router.Group("/")
	{
		public.GET("/", Hello)
		public.GET("/sign-up", SignUp)
		public.POST("/sign-up", ProcessSignUp)
		public.POST("/login", Login)
		public.GET("/automatic-redirect", AutomaticRedirect)
	}

	protected := router.Group("/", AuthRequired)
	{
		protected.GET("/dashboard", Dashboard)
	}

	// Listen and server on 0.0.0.0:8080
	err := router.Run(":" + config.Port())

	utils.Log("Dingus stopping")

	if err != nil {
		utils.LogError(err)
	}
}

func AuthRequired(c *gin.Context) {

	// before request, make sure we have a session
	// and that they have a handle
	session := sessions.Get(c)

	if !session.IsLoggedIn() {
		utils.Log("Nobody logged in, going to login screen.")
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}

	c.Next()

	// after request
}

func Dashboard(c *gin.Context) {
	session := sessions.Get(c)
	content := utils.Render("templates/dashboard.tmpl", gin.H{"handle": session.Handle, "client_id": config.ClientId()})
	c.HTML(http.StatusOK, "master.tmpl", gin.H{"content": template.HTML(content)})
}

func Hello(c *gin.Context) {

	session := sessions.Get(c)
	if session.IsLoggedIn() {
		c.Redirect(http.StatusMovedPermanently, "/dashboard")
		return
	}

	content := utils.Render("templates/index.tmpl", gin.H{"handle": session.Handle, "loggedIn": false})
	c.HTML(http.StatusOK, "master.tmpl", gin.H{"content": template.HTML(content)})
}

func SignUp(c *gin.Context) {

	c.HTML(http.StatusOK, "sign-up.tmpl", gin.H{})
}

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
