package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Session struct {
	session sessions.Session
	options sessions.Options
}

type Options struct {
	sessions.Options
}

func NewSession(c *gin.Context) *Session {
	session := sessions.Default(c)
	return &Session{session: session}
}

func (s *Session) Get(key string) interface{} {
	return s.session.Get(key)
}

func (s *Session) Set(key string, value interface{}) {
	s.session.Set(key, value)
}

func (s *Session) Save() error {
	return s.session.Save()
}

func (s *Session) Disable() error {
	s.session.Clear()
	return s.Save()
}