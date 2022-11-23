package community

import (
	"encoding/json"
	"net/http"

	"github.com/mariobac1/api_/domain/community"
	"github.com/mariobac1/api_/models"
)

type handler struct {
	usecase community.Storage
}

func newHandler(usecase community.Storage) handler {
	return handler{usecase}
}

func (h *handler) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "message": "Method not permit"}`))
		return
	}
	data := models.Community{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "message": "The structure is wrong"}`))
		return
	}
	err = h.usecase.Create(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type": "error", "message": "An issue occurs when try create a communities "}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message_type": "message", "message": "Community created Ok"}`))

}
