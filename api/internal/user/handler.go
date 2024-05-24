package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Handler struct {
	userService Service
}

func NewHandler(repository Repository) Handler {
	return Handler{userService: NewService(repository)}
}

func (h Handler) RegisterUserRouteHandlers(mux *http.ServeMux) {
	mux.HandleFunc("POST /user", h.registerUserHandler)
}

func (h Handler) registerUserHandler(response http.ResponseWriter, request *http.Request) {
	var registerUser RegisterUser

	err := json.NewDecoder(request.Body).Decode(&registerUser)
	if err != nil {
		http.Error(response, "invalid request", http.StatusBadRequest)
		return
	}

	err = h.userService.RegisterUser(registerUser)

	if errors.Is(err, &InvalidUserError{}) {
		http.Error(response, "invalid request", http.StatusBadRequest)
		return
	} else if errors.Is(err, &UserExistsError{}) {
		http.Error(response, "user with email exists", http.StatusConflict)
		return
	} else if err != nil {
		fmt.Println(err.Error())
		http.Error(response, "internal server error", http.StatusInternalServerError)
		return
	}

	if registerUser.isMissingInformation() {
		http.Error(response, "invalid request", http.StatusBadRequest)
		return
	}
}
