package user

import (
	"fmt"
	"net/http"

	"github.com/HenryGunadi/simple-chat-app/server/services/auth"
	"github.com/HenryGunadi/simple-chat-app/server/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	authService *auth.AuthService
}

func NewUserHandler(authService *auth.AuthService) *Handler {
	return &Handler{authService: authService}
} 

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/user", auth.RequireAuth(h.HandleSendUserData, h.authService)).Methods("GET")
}

func (h *Handler) HandleSendUserData(w http.ResponseWriter, r *http.Request) {
	u, err := h.authService.GetUserSession(r)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error getting user data from cookies : %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, u)
} 


