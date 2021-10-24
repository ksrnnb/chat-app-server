package middleware

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewSessionMiddleware() gin.HandlerFunc {
	store := cookie.NewStore([]byte(os.Getenv("SESSION_KEY")))
	store.Options(options())
	return sessions.Sessions("sessionId", store)
}

func options() sessions.Options {
	isSecure := os.Getenv("APP_ENV") == "production"

	return sessions.Options{
		Secure:   isSecure,
		HttpOnly: true,
	}
}
