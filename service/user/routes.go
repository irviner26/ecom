package user

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/irviner26/ecom/config"
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
	router.POST("/api/v1/auth/login", h.handleLogin)
	router.POST("/api/v1/auth/register", h.handleRegister)
}

func (h *Handler) handleLogin(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// get JSON payload
	var payload types.LoginUserPayload
	if err := utils.ParseJson(request, &payload); err != nil {
		utils.WriteError(writer, http.StatusBadRequest, err)
	}

	// validate the payload
	if err := utils.Validate.StructCtx(request.Context(), payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(writer, http.StatusBadRequest, errors)
	}

	// get the user from storage
	user, err := h.store.GetUserByEmail(payload.Email, request.Context())
	if err != nil {
		utils.WriteError(
			writer,
			http.StatusNotFound,
			errors.New("User does not exist"),
		)

		return
	}

	// match the password
	if !auth.IsCorrectPassword([]byte(user.Password), []byte(payload.Password)) {
		utils.WriteError(
			writer,
			http.StatusUnauthorized,
			errors.New("Wrong password"),
		)

		return
	}

	secret := []byte(config.Global.JWTSecret)
	authToken, err := auth.CreateJWT(secret, user.ID)
	if err != nil {
		utils.WriteError(
			writer,
			http.StatusInternalServerError,
			err,
		)

		return
	}

	utils.WriteJson(
		writer,
		http.StatusOK,
		map[string]string{
			"token": authToken,
		},
	)
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
	if err == nil {
		utils.WriteError(
			writer,
			http.StatusConflict,
			errors.New("user with email '" + payload.Email + "' already exists"),
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
	}, request.Context())
	if err != nil {
		utils.WriteError(writer, http.StatusInternalServerError, err)

		return
	}

	// write success user creation response
	utils.WriteJson(writer, http.StatusCreated, nil)
}
