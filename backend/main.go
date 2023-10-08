package main

import (
	"backend/handler"
	"backend/service"
	"log"
	"net/http"
)

func main() {
	nftService, err := service.NewNFTService()
	if err != nil {
		log.Fatalf("Failed to create NFT service: %v", err)
	}
	auctionService, err := service.NewAuctionService()
	if err != nil {
		log.Fatalf("Failed to create Auction service: %v", err)
	}
	ownershipService, err := service.NewOwnershipService()
	if err != nil {
		log.Fatalf("Failed to create Ownership service: %v", err)
	}

	nftHandler := handler.NewNFTHandler(nftService)
	auctionHandler := handler.NewAuctionHandler(auctionService)
	ownershipHandler := handler.NewOwnershipHandler(ownershipService)

	router := NewRouter(nftHandler, auctionHandler, ownershipHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
