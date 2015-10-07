package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

	logger := createLogger()
	logger.Println("Started")

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", hello)
	router.GET("/automatic-redirect", AutomaticRedirect)
	router.Static("/public", "./public")

	// Listen and server on 0.0.0.0:8080
	err := router.Run(":8080")

	logger.Println("Stopping")

	if err != nil {
		logger.Println(err.Error())
	}
}

func hello(c *gin.Context) {

	content := Render("templates/index.tmpl", gin.H{"name": "Dingus"})
	c.HTML(http.StatusOK, "master.tmpl", gin.H{"content": template.HTML(content)})
}

func Render(fileName string, data map[string]interface{}) string {
	t, err := template.ParseFiles(fileName)
	buff := bytes.NewBufferString("")
	if err == nil {
		err := t.Execute(buff, data)
		if err == nil {
			return buff.String()
		}
	}

	panic(err)
}

func AutomaticRedirect(c *gin.Context) {

	oauthUrl := "https://accounts.automatic.com/oauth/access_token"
	postData := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {c.Query("code")},
		"client_id":     {ClientId()},
		"client_secret": {ClientSecret()},
	}

	resp, _ := http.PostForm(oauthUrl, postData)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	c.HTML(http.StatusOK, "master.tmpl", gin.H{"content": string(body)})
}
