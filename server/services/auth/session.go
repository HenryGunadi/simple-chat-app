package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
)

const (
	SessionName = "user"
)


type SessionOptions struct {
	CookiesKey string
	MaxAge     int
	HttpOnly   bool
	Secure     bool
}

func NewCookieStore(opt SessionOptions) *sessions.CookieStore {
	store := sessions.NewCookieStore([]byte(opt.CookiesKey))

	store.MaxAge(opt.MaxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = opt.HttpOnly
	store.Options.Secure = opt.Secure
	store.Options.SameSite = http.SameSiteLaxMode

	return store
}