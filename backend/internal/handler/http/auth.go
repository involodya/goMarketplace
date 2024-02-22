package handler

import (
	"fullstack/backend/internal/entity"
	"net/http"
	"strings"
)

var AuthHeaderName = "Authorization"

func (h *Handler) getClaimsFromAuthHeader(r *http.Request) (*map[string]string, error) {
	jwtClaims := &map[string]string{}

	authHeader := r.Header.Get(AuthHeaderName)
	if authHeader == "" {
		return jwtClaims, entity.ErrEmptyAuthHeader
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return jwtClaims, entity.ErrInvalidAuthHeader
	}

	jwtClaims, err := h.auth.FetchAuthn(headerParts[1])
	return jwtClaims, err
}
