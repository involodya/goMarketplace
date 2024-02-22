package handler

import (
	"encoding/json"
	"fullstack/backend/internal/entity"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) getUserItems(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		authClaims, _ := h.getClaimsFromAuthHeader(r)
		authUserID, _ := strconv.ParseUint((*authClaims)["sub"], 10, 32)
		log.Println("Getting user items for user with id", authUserID)

		items, err := h.service.Item.GetBySeller(uint(authUserID))
		if err != nil {
			log.Println("Error occurred in handler.getItemByID, Error:", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(items)
		if err != nil {
			log.Printf("Error when encoding users, Error: %v\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) getItemInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		vars := mux.Vars(r)
		itemID, err := strconv.ParseUint(vars["item_id"], 10, 32)
		if err != nil {
			log.Printf("Error when parsing item_id to uint in handler.getItemInfo, Error: %v", err.Error())
			return
		}

		item, err := h.service.Item.Get(uint(itemID))
		if err != nil {
			log.Println("Error occurred in handler.getItemByID, Error:", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		authClaims, _ := h.getClaimsFromAuthHeader(r)
		authUserID, _ := strconv.ParseUint((*authClaims)["sub"], 10, 32)
		if item.SellerID != uint(authUserID) {
			log.Printf("Access denied, UserID: %d, authUserID: %d", item.SellerID, authUserID)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(item)
		if err != nil {
			log.Printf("Error when encoding item, Error: %v\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) updateItemInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		authClaims, _ := h.getClaimsFromAuthHeader(r)
		authUserID, _ := strconv.ParseUint((*authClaims)["sub"], 10, 32)

		vars := mux.Vars(r)
		itemID, err := strconv.ParseUint(vars["item_id"], 10, 32)
		if err != nil {
			log.Printf("Error when parsing item_id to uint in handler.updateItemInfo, Error: %v", err.Error())
			return
		}

		var item entity.Item
		err = json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			log.Printf("Error when decoding request body to item in handler.updateItemInfo, Error: %v", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		itemDb, err := h.service.Item.Get(uint(itemID))
		if err != nil {
			log.Println("Error occurred in handler.getItemByID, Error:", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if itemDb.SellerID != uint(authUserID) {
			log.Printf("Access denied, UserID: %d, authUserID: %d", itemDb.SellerID, authUserID)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		item.ID = uint(itemID)
		item.SellerID = uint(authUserID)
		err = h.service.Item.Update(&item)
		if err != nil {
			log.Printf("Error when updating item in handler.updateItemInfo, Error: %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(item)
		if err != nil {
			log.Printf("Error when encoding item, Error: %v\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) createItemInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		authClaims, _ := h.getClaimsFromAuthHeader(r)
		authUserID, _ := strconv.ParseUint((*authClaims)["sub"], 10, 32)

		var item entity.Item
		err := json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			log.Printf("Error when decoding request body to item in handler.createItemInfo, Error: %v", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		item.SellerID = uint(authUserID)
		err = h.service.Item.Create(&item)
		if err != nil {
			log.Printf("Error when updating item in handler.createItemInfo, Error: %v", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(item)
		if err != nil {
			log.Printf("Error when encoding item, Error: %v\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
