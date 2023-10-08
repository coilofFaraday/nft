package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"time"
)

// Auction represents an NFT auction
type Auction struct {
	NFT           string    `json:"nft"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	StartPrice    int       `json:"start_price"`
	HighestBid    int       `json:"highest_bid"`
	HighestBidder string    `json:"highest_bidder"`
}

// StartAuction initiates an auction for an NFT
func (s *SmartContract) StartAuction(ctx contractapi.TransactionContextInterface, nft string, startTime time.Time, endTime time.Time, startPrice int) error {
	auction := Auction{
		NFT:        nft,
		StartTime:  startTime,
		EndTime:    endTime,
		StartPrice: startPrice,
	}

	auctionBytes, _ := json.Marshal(auction)

	return ctx.GetStub().PutState(nft, auctionBytes)
}

// Bid places a new bid on an NFT
func (s *SmartContract) Bid(ctx contractapi.TransactionContextInterface, nft string, bid int, bidder string) error {
	auctionBytes, _ := ctx.GetStub().GetState(nft)

	var auction Auction
	_ = json.Unmarshal(auctionBytes, &auction)

	if bid > auction.HighestBid {
		auction.HighestBid = bid
		auction.HighestBidder = bidder
	}

	auctionBytes, _ = json.Marshal(auction)

	return ctx.GetStub().PutState(nft, auctionBytes)
}
