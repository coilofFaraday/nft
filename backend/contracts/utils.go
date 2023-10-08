package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"time"
)

// CheckAuction checks if an auction is valid
func CheckAuction(ctx contractapi.TransactionContextInterface, nft string) (*Auction, error) {
	auctionBytes, err := ctx.GetStub().GetState(nft)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state: %v", err)
	}

	var auction Auction
	_ = json.Unmarshal(auctionBytes, &auction)

	if time.Now().Before(auction.StartTime) {
		return nil, fmt.Errorf("Auction has not started yet")
	}

	if time.Now().After(auction.EndTime) {
		return nil, fmt.Errorf("Auction has already ended")
	}

	return &auction, nil
}

// CheckOwner checks if a user is the owner of an NFT
func CheckOwner(ctx contractapi.TransactionContextInterface, nft string, owner string) (bool, error) {
	nftBytes, err := ctx.GetStub().GetState(nft)

	if err != nil {
		return false, fmt.Errorf("Failed to read from world state: %v", err)
	}

	var nftAsset NFT
	_ = json.Unmarshal(nftBytes, &nftAsset)

	if nftAsset.Owner != owner {
		return false, nil
	}

	return true, nil
}
