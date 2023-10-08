package handler

import (
	"backend/model"
	"backend/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type AuctionHandler struct {
	Service *service.AuctionService
}

func NewAuctionHandler(service *service.AuctionService) *AuctionHandler {
	return &AuctionHandler{
		Service: service,
	}
}

func (handler *AuctionHandler) StartAuctionHandler(w http.ResponseWriter, r *http.Request) {
	var auction model.Auction
	_ = json.NewDecoder(r.Body).Decode(&auction)

	result, err := handler.Service.StartAuction(auction.NFT, auction.StartTime, auction.EndTime, auction.StartPrice)
	if err != nil {
		http.Error(w, "Failed to start auction", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (handler *AuctionHandler) BidHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nft := vars["nft"]

	var bid struct {
		Bid    int    `json:"bid"`
		Bidder string `json:"bidder"`
	}
	_ = json.NewDecoder(r.Body).Decode(&bid)

	result, err := handler.Service.Bid(nft, bid.Bid, bid.Bidder)
	if err != nil {
		http.Error(w, "Failed to place bid", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}
