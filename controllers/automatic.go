package controllers

import (
	"github.com/benreic/dingus/config"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
)

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
