package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) getUserStats(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		authClaims, _ := h.getClaimsFromAuthHeader(r)
		authUserID, _ := strconv.ParseUint((*authClaims)["sub"], 10, 32)

		orders, err := h.service.Order.GetBySeller(uint(authUserID))
		if err != nil {
			log.Println("Error occurred in handler.getItemByID, Error:", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		itemCount := 0
		for _, order := range *orders {
			itemCount += int(order.Count)
		}

		stats := map[string]interface{}{
			"item_count": itemCount,
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(stats)
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
