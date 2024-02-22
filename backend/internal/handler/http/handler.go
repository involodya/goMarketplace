package handler

import (
	"fmt"
	"fullstack/backend/internal/entity"
	"fullstack/backend/internal/pkg/auth"
	"fullstack/backend/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Handler struct {
	service *service.Service
	auth    entity.AuthManager
}

func NewHandler(service *service.Service, auth *auth.AuthManager) *Handler {
	return &Handler{
		service: service,
		auth:    auth,
	}
}

func (h *Handler) NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", index).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/user", h.getUserInfo).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/register", h.registerUser).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/login", h.loginUser).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/stats", h.getUserStats).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/items", h.getUserItems).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/item/{item_id}", h.getItemInfo).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/item/create", h.createItemInfo).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/item/{item_id}", h.updateItemInfo).Methods(http.MethodPost, http.MethodOptions)

	r.Use(LoggingMiddleware(r))
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(CustomCORSMiddleware(r))

	return r
}

func LoggingMiddleware(r *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			log.Printf("Origin: %s | Forwarded: %s | Method: %s | RequestURI: %s", req.Header.Get("Origin"), req.Header.Get("Forwarded"), req.Method, req.RequestURI)

			next.ServeHTTP(w, req)
		})
	}
}

func CustomCORSMiddleware(r *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Length, Content-Type, Authorization, Host, Origin, X-CSRF-Token")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization")

			next.ServeHTTP(w, req)
		})
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!\n")
}
