package handler

import (
	"clevergo.tech/jsend"
	"encoding/json"
	"fullstack/backend/internal/entity"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) loginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var userLogin entity.UserLogin
		err := json.NewDecoder(r.Body).Decode(&userLogin)
		if err != nil {
			log.Println("Error occurred in handler.loginUser when decoding userLogin:", err.Error())
			w.WriteHeader(http.StatusTeapot)
			return
		}

		userId, err := h.service.User.Login(&userLogin)
		if err != nil {
			log.Println("Error occurred in handler.loginUser when logging user:", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			jsend.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tokenString, _ := h.auth.MakeAuthn(userId)

		w.Header().Add(AuthHeaderName, "Bearer "+tokenString)

		log.Println("Login as user with id ", userId)
		w.WriteHeader(http.StatusOK)
		jsend.Success(w, http.StatusOK)
	} else if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) registerUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var userReg entity.UserRegister
		err := json.NewDecoder(r.Body).Decode(&userReg)
		if err != nil {
			log.Println("Error occurred in handler.registerUser when decoding userReg:", err.Error())
			w.WriteHeader(http.StatusTeapot)
			jsend.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = h.service.User.Register(&userReg)
		if err != nil {
			log.Println("Error occurred in handler.registerUser when registering user:", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			jsend.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println("Registered user with email", userReg.Email)
		w.WriteHeader(http.StatusCreated)
		json_msg := map[string]string{
			"message": "successful registration",
		}
		jsend.Success(w, json_msg, http.StatusCreated)
	} else if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) getUserInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		authClaims, _ := h.getClaimsFromAuthHeader(r)
		authUserID, _ := strconv.ParseUint((*authClaims)["sub"], 10, 32)

		userInfo, err := h.service.User.Get(uint(authUserID))
		if err != nil {
			log.Printf("Error occurred in handler.getUserInfo, Error: %v", err.Error())
			return
		}

		err = json.NewEncoder(w).Encode(userInfo)
		if err != nil {
			log.Printf("Error when encoding userInfo, Error: %v\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}

		log.Println("Get user info email:", userInfo.Email)
	} else if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
