package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/HenryGunadi/simple-chat-app/server/config"
	"github.com/HenryGunadi/simple-chat-app/server/services/auth"
	"github.com/HenryGunadi/simple-chat-app/server/services/user"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
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
	AuthStore := auth.NewAuthStore(s.db)
	AuthHandler := auth.NewAuthHandler(AuthStore, AuthService)
	AuthHandler.RegisteredRoutes(router)

	// user service
	userStore := user.NewStore(s.db)
	userHandler := user.NewUserHandler(userStore, AuthService)
	userHandler.RegisterRoutes(router)

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