package handler

import (
	"backend/model"
	"backend/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type NFTHandler struct {
	Service *service.NFTService
}

func NewNFTHandler(service *service.NFTService) *NFTHandler {
	return &NFTHandler{
		Service: service,
	}
}

func (handler *NFTHandler) CreateNFTHandler(w http.ResponseWriter, r *http.Request) {
	var nft model.NFT
	_ = json.NewDecoder(r.Body).Decode(&nft)

	result, err := handler.Service.CreateNFT(nft.Name, nft.Unique, nft.Owner)
	if err != nil {
		http.Error(w, "Failed to create NFT", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (handler *NFTHandler) QueryNFTHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	nft, err := handler.Service.QueryNFT(name)
	if err != nil {
		http.Error(w, "Failed to query NFT", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(nft)
}
