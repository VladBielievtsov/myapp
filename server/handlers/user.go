package handlers

import (
	"encoding/json"
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
