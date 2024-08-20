package api

import (
	"log"
	"net/http"

	"github.com/HenryGunadi/simple-chat-app/server/config"
	"github.com/HenryGunadi/simple-chat-app/server/services/auth"
	"github.com/HenryGunadi/simple-chat-app/server/services/chat"
	"github.com/HenryGunadi/simple-chat-app/server/services/user"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{addr: addr}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	// auth service
	sessionStore := auth.NewCookieStore(auth.SessionOptions{
		CookiesKey: config.Envs.CookiesAuthSecret,
		MaxAge: config.Envs.CookiesAuthAgeinSeconds,
		HttpOnly: config.Envs.CookiesAuthIsHttpOnly,
		Secure: config.Envs.CookiesAuthIsSecure,
	})
	AuthService := auth.NewAuthService(sessionStore)
	AuthHandler := auth.NewAuthHandler(AuthService)
	AuthHandler.RegisteredRoutes(router)

	// user service
	userHandler := user.NewUserHandler(AuthService)
	userHandler.RegisterRoutes(router)

	// websocker service
	hub := chat.NewHub()
	go hub.Run()
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(hub, w, r)
	})

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("../client/public"))))

	// cors
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
        AllowedHeaders:   []string{"Authorization", "Content-Type"},
        AllowCredentials: true,
	})
	corsHandler := c.Handler(router)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, corsHandler)
}