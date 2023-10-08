package handler

import (
	"backend/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type OwnershipHandler struct {
	Service *service.OwnershipService
}

func NewOwnershipHandler(service *service.OwnershipService) *OwnershipHandler {
	return &OwnershipHandler{
		Service: service,
	}
}

func (handler *OwnershipHandler) TransferOwnershipHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nft := vars["nft"]

	var owner struct {
		Owner string `json:"owner"`
	}
	_ = json.NewDecoder(r.Body).Decode(&owner)

	result, err := handler.Service.TransferOwnership(nft, owner.Owner)
	if err != nil {
		http.Error(w, "Failed to transfer ownership", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}
