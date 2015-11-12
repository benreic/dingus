package middleware

import (
	"github.com/benreic/dingus/sessions"
	"github.com/benreic/dingus/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
