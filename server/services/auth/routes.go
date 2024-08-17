package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HenryGunadi/simple-chat-app/server/types"
	"github.com/gorilla/mux"
	"github.com/markbates/goth/gothic"
)

type Handler struct {
	store types.AuthStore
	authService *AuthService
}

func NewAuthHandler(store types.AuthStore, authService *AuthService) *Handler {
	return &Handler{store: store, authService: authService}
}

func (h *Handler) RegisteredRoutes(r *mux.Router) {
	r.HandleFunc("/auth/{provider}", h.HandleProviderLogin).Methods("GET")
	r.HandleFunc("/auth/{provider}/callback", h.HandleAuthCallbackFunction).Methods("GET")
}

func (h *Handler) HandleProviderLogin(w http.ResponseWriter, r *http.Request) {
	if u, err := gothic.CompleteUserAuth(w, r); err == nil {
		log.Println("User is authenticated : ", u)
		http.Redirect(w, r, "http://localhost:3000", http.StatusFound)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

func (h *Handler) HandleAuthCallbackFunction(w http.ResponseWriter, r *http.Request) {
	u, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	log.Printf("user name : %v", u.Name)

	err = h.authService.StoreUserSession(w, r, u)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Redirecting to home")


	http.Redirect(w, r, "http://localhost:3000", http.StatusTemporaryRedirect)
}

