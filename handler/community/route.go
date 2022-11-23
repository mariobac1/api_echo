package community

import (
	"net/http"

	"github.com/mariobac1/api_/domain/community"
)

func RouteCommunity(mux *http.ServeMux, usecase community.Storage) {
	h := newHandler(usecase)

	mux.HandleFunc("/v1/communities/create", h.create)
}
