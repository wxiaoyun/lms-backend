package session

import (
	"encoding/gob"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	Store *session.Store
)

const (
	CookieKey = "token"
	MaxAge    = time.Hour * 24 * 7 // 7 days
)

func SetupStore() {
	Store = session.New(session.Config{
		Expiration:     MaxAge,
		CookieHTTPOnly: true,
	})

	gob.Register(LoginSession{})
}
