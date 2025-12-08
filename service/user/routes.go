package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler struct {

}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *httprouter.Router) {
	router.POST("/api/v1/login", h.handleLogin)
	router.POST("/api/v1/register", h.handleRegister)
}

func (h *Handler) handleLogin(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	
}

func (h *Handler) handleRegister(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	
}