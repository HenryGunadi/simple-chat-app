package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HenryGunadi/simple-chat-app/server/config"
	"github.com/HenryGunadi/simple-chat-app/server/utils"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
)

type AuthService struct {
}

func NewAuthService(store sessions.Store) *AuthService {
	gothic.Store = store

	goth.UseProviders(
		discord.New(
			config.Envs.DCClientID,
			config.Envs.DCSecret,
			buildCallbackURL("discord"),
		),
	)

	return &AuthService{}
}

func (s *AuthService) GetUserSession(r *http.Request) (goth.User, error) {
	session, err := gothic.Store.Get(r, SessionName)
	if err != nil {
		return goth.User{}, err
	}

	u := session.Values["user"]
	if u == nil {
		return goth.User{}, fmt.Errorf("user is not authenticated! %v", u)
	}

	return u.(goth.User), nil
} 

func RequireAuth(handlerFunc http.HandlerFunc, auth *AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := auth.GetUserSession(r)
		if err != nil {
			log.Println("User is not authorized!")
			utils.WriteError(w, http.StatusUnauthorized, err)
			return
		}

		log.Printf("User is authenticated : %v", u.Name)
		log.Printf("user data : %v", u)

		handlerFunc(w, r)
	}
}

func (s *AuthService) StoreUserSession(w http.ResponseWriter, r *http.Request, user goth.User) error {
	session, _ := gothic.Store.Get(r, SessionName)
	
	session.Values["user"] = user

	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}


func buildCallbackURL(provider string) string {
	return fmt.Sprintf("http://localhost:8080/auth/%s/callback", provider)
}