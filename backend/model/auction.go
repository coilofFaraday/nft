package model

import "time"

type Auction struct {
	NFT           string    `json:"nft"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	StartPrice    int       `json:"start_price"`
	HighestBid    int       `json:"highest_bid"`
	HighestBidder string    `json:"highest_bidder"`
}
