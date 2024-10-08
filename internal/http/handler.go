package http

import (
	"github.com/gorilla/mux"
	"github.com/my-otp/otp-rest/internal/service/ports"
)

type Handler struct {
	router *mux.Router
	db     ports.Database
}

func NewHandler() *Handler {
	router := mux.NewRouter()
	router.Use()
	return &Handler{
		router: router,
	}
}
