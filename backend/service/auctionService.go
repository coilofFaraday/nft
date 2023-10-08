package service

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"strconv"
	"time"
)

type AuctionService struct {
	Contract *gateway.Contract
}

func NewAuctionService() (*AuctionService, error) {
	// This would be similar to NewNFTService. You would just replace "nftcc" with the appropriate chaincode name if different.
	cp, err := gateway.FromConfig("/contracts/connection.yaml")
	if err != nil {
		return nil, err
	}

	gw, err := gateway.Connect(
		gateway.WithConfig(cp),
		gateway.WithIdentity(gateway.NewX509Identity("Org1MSP", "/contracts/userid.pem")),
	)
	if err != nil {
		return nil, err
	}

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		return nil, err
	}

	contract := network.GetContract("nftcc")

	return &AuctionService{
		Contract: contract,
	}, nil
}

func (service *AuctionService) StartAuction(nft string, startTime time.Time, endTime time.Time, startPrice int) (string, error) {
	result, err := service.Contract.SubmitTransaction("StartAuction", nft, startTime.String(), endTime.String(), strconv.Itoa(startPrice))
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func (service *AuctionService) Bid(nft string, bid int, bidder string) (string, error) {
	result, err := service.Contract.SubmitTransaction("Bid", nft, strconv.Itoa(bid), bidder)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
