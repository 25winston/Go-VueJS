package sessions

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/kataras/iris/sessions"
)

// NewSessionCookie :
func NewSessionCookie() *sessions.Sessions {

	expires, err := strconv.ParseInt(os.Getenv("SESSION_EXPIRES"), 10, 64)
	if err != nil {
		expires = 60
	}

	var config sessions.Config

	if strings.ToLower(os.Getenv("SESSION_SECURE")) == "true" {
		hash := []byte(os.Getenv("SESSION_SECURE_HASH"))
		secret := []byte(os.Getenv("SESSION_SECURE_SECRET"))
		secure := securecookie.New(hash, secret)

		config = sessions.Config{
			Cookie:       os.Getenv("SESSION_NAME"),
			Expires:      time.Duration(expires) * time.Minute,
			Encode:       secure.Encode,
			Decode:       secure.Decode,
			AllowReclaim: true,
		}

	} else {
		config = sessions.Config{
			Cookie:       os.Getenv("SESSION_NAME"),
			Expires:      time.Duration(expires) * time.Minute,
			AllowReclaim: true,
		}
	}

	SessionManager := sessions.New(config)

	return SessionManager
}
