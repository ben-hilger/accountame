package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Handler struct {
	accountService accountService
}

func NewHandler(repository repository) Handler {
	return Handler{accountService: newService(repository)}
}

func (h Handler) LoginUserHandler(response http.ResponseWriter, request *http.Request) {
	var loginRequest loginAccountRequest

	err := json.NewDecoder(request.Body).Decode(&loginRequest)

	accountId, err := h.accountService.LoginUser(loginRequest)
	if errors.Is(err, &existsError{}) || errors.Is(err, &invalidPasswordError{}) || errors.Is(err, &invalidAccountError{}) {
		http.Error(response, "invalid request", http.StatusBadRequest)
		return
	} else if err != nil {
		fmt.Println(err)
		http.Error(response, "internal server error", http.StatusInternalServerError)
		return
	}

	tokenString, err := h.accountService.CreateAuthenticationToken(accountId)
	if err != nil {
		http.Error(response, "internal server error", http.StatusInternalServerError)
		return
	}

	loginResponse := loginAccountResponse{AuthToken: tokenString}
	err = json.NewEncoder(response).Encode(loginResponse)
	if err != nil {
		http.Error(response, "internal server error", http.StatusInternalServerError)
	}
}

func (h Handler) GetUserInformationHandler(response http.ResponseWriter, request *http.Request) {
	uid, ok := request.Context().Value("userId").(string)
	if !ok || uid == "" {
		http.Error(response, "invalid request", http.StatusBadRequest)
		return
	}
	account, err := h.accountService.GetUser(uid)
	if errors.Is(err, &existsError{}) {
		http.Error(response, "invalid request", http.StatusBadRequest)
		return
	} else if err != nil {
		fmt.Println(err)
		http.Error(response, "internal server error", http.StatusInternalServerError)
		return
	}

	type accountInformationResponse struct {
		Id       string `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
		Name     string `json:"name"`
	}

	informationResponse := accountInformationResponse{
		Id:       account.Id,
		Email:    account.Email,
		Username: account.Username,
		Name:     account.Name,
	}
	err = json.NewEncoder(response).Encode(informationResponse)
	if err != nil {
		http.Error(response, "internal server error", http.StatusInternalServerError)
	}

}

func (h Handler) RegisterUserHandler(response http.ResponseWriter, request *http.Request) {
	var registerUser registerAccount

	err := json.NewDecoder(request.Body).Decode(&registerUser)
	if err != nil {
		http.Error(response, "invalid request", http.StatusBadRequest)
		return
	}

	err = h.accountService.RegisterUser(registerUser)

	if errors.Is(err, &invalidAccountError{}) {
		http.Error(response, "invalid request", http.StatusBadRequest)
		return
	} else if errors.Is(err, &existsError{}) {
		http.Error(response, "Account with email exists", http.StatusConflict)
		return
	} else if err != nil {
		fmt.Println(err.Error())
		http.Error(response, "internal server error", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
}
