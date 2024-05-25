package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"my-app/middlewares"
	"my-app/services"
	"my-app/types"
	"my-app/utils"
	"net/http"
)

var userService = services.NewUserService()

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := userService.GetUsers()
	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, err)
		return
	}

	utils.JSONResponse(w, http.StatusOK, types.FilterUsers(users))
}

func StoreUser(w http.ResponseWriter, r *http.Request) {
	var req types.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"status": "fail", "message": "Invalid request payload"})
		return
	}

	user, err := userService.CreateUser(req)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSONResponse(w, http.StatusCreated, types.FilterUser(&user))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := userService.GetUserByID(id)
	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, err)
		return
	}

	utils.JSONResponse(w, http.StatusOK, user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req types.LoginUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"status": "fail", "message": "Invalid request payload"})
		return
	}

	user, tokenString, err := userService.Login(req)
	if err != nil {
		utils.JSONResponse(w, http.StatusUnauthorized, err)
		return
	}

	result := map[string]interface{}{
		"token": tokenString,
		"user":  types.FilterUser(&user),
	}

	utils.JSONResponse(w, http.StatusOK, result)
}

func GetMe(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middlewares.AuthContext{}).(types.User)
	if !ok {
		utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"message": "Could not retrieve user from context"})
		return
	}

	utils.JSONResponse(w, http.StatusOK, types.FilterUser(&user))
}
