package main

import (
	"backend/handler"
	"github.com/gorilla/mux"
)

func NewRouter(nftHandler *handler.NFTHandler, auctionHandler *handler.AuctionHandler, ownershipHandler *handler.OwnershipHandler) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/nft/create", nftHandler.CreateNFTHandler).Methods("POST")
	router.HandleFunc("/nft/{name}", nftHandler.QueryNFTHandler).Methods("GET")

	router.HandleFunc("/auction/start", auctionHandler.StartAuctionHandler).Methods("POST")
	router.HandleFunc("/auction/{nft}/bid", auctionHandler.BidHandler).Methods("POST")

	router.HandleFunc("/ownership/{nft}/transfer", ownershipHandler.TransferOwnershipHandler).Methods("POST")

	return router
}
