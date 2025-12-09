package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/irviner26/ecom/service/auth"
	"github.com/irviner26/ecom/types"
	"github.com/irviner26/ecom/utils"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *httprouter.Router) {
	router.POST("/api/v1/login", h.handleLogin)
	router.POST("/api/v1/register", h.handleRegister)
}

func (h *Handler) handleLogin(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func (h *Handler) handleRegister(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// get JSON payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJson(request, &payload); err != nil {
		utils.WriteError(writer, http.StatusBadRequest, err)
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(writer, http.StatusBadRequest, errors)

		return
	}

	// check if user already exists
	_, err := h.store.GetUserByEmail(payload.Email, request.Context())
	if err != nil {
		utils.WriteError(
			writer,
			http.StatusConflict,
			fmt.Errorf("User with email %s already exists", payload.Email),
		)

		return
	}

	// hash user's password
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(writer, http.StatusInternalServerError, err)

		return
	}

	// store new user to database
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(writer, http.StatusInternalServerError, err)

		return
	}

	// write success user creation response
	utils.WriteJson(writer, http.StatusCreated, nil)
}
