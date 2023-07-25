package session

import "github.com/gofiber/fiber/v2"

type LoginSession struct {
	UserID         uint
	Email          string
	IsMasquerading bool
}

func GetLoginSession(c *fiber.Ctx) (LoginSession, error) {
	sess, err := Store.Get(c)
	if err != nil {
		return LoginSession{}, err
	}

	loginSession := sess.Get(CookieKey)
	if loginSession == nil {
		return LoginSession{}, nil
	}

	//nolint:forcetypeassert // we know that loginSession is a LoginSession
	return loginSession.(LoginSession), nil
}
