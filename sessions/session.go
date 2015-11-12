package sessions

import (
	"github.com/benreic/dingus/config"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var sessionName = "dingus-session"

/**
 * The DingusSession struct that holds all of
 * our strongly typed session bits.
 *
 * @author Ben Reichelt <benr@clockwork.net>
 *
**/

type DingusSession struct {
	Handle   string
	DriverId int
}

func GetCookieStore() *sessions.CookieStore {
	return sessions.NewCookieStore([]byte(config.SessionAuthenticationKey()), []byte(config.SessionEncryptionKey()))
}

/**
 * Gets the session for the current request
 *
 * @author Ben Reichelt <benr@clockwork.net>
 *
 * @param   c      gin.Context
 * @return  DingusSession
**/

func Get(c *gin.Context) *DingusSession {

	store := GetCookieStore()
	session, _ := store.Get(c.Request, sessionName)
	dingusSession := new(DingusSession)

	if session == nil || session.Values["handle"] == nil {
		dingusSession.Handle = ""
		dingusSession.DriverId = 0
		return dingusSession
	}

	handle := session.Values["handle"].(string)
	if len(handle) == 0 {
		dingusSession.Handle = ""
		dingusSession.DriverId = 0
		return dingusSession
	}

	dingusSession.Handle = handle
	driverId := session.Values["driverId"].(int)
	dingusSession.DriverId = driverId
	return dingusSession
}

/**
 * For the current DingusSession save, the values to the
 * session storage
 *
 * @author Ben Reichelt <benr@clockwork.net>
 *
 * @param   c      gin.Context
 * @return  void
**/

func (s *DingusSession) Save(c *gin.Context) {

	store := GetCookieStore()
	session, _ := store.Get(c.Request, sessionName)
	session.Values["handle"] = s.Handle
	session.Values["driverId"] = s.DriverId
	session.Save(c.Request, c.Writer)
}

/**
 * Determines if the current session is logged in or not
 *
 * @author Ben Reichelt <benr@clockwork.net>
 *
 * @return  bool
**/

func (s *DingusSession) IsLoggedIn() bool {
	return s.Handle != "" && len(s.Handle) > 0
}
