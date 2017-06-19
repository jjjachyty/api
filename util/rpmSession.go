package util

import (
	"github.com/tango-contrib/session"
	"time"
)

var defaultSession *session.Sessions

type RpmSession struct {
}

func (r RpmSession) GetDefultSesion() *session.Session {
	return defaultSession.SessionFromID(session.DefaultSessionIdName)
}

func init() {
	defaultSession = session.New(session.Options{
		MaxAge: time.Minute * 10,
	})
}
