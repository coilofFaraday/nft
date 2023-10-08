package service

import (
	"backend/model"
	"encoding/json"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type NFTService struct {
	Contract *gateway.Contract
}

func NewNFTService() (*NFTService, error) {
	// Load connection profile; it has details about fabric network
	cp, err := gateway.FromConfig("/contracts/connection.yaml")
	if err != nil {
		return nil, err
	}

	// Connect to gateway peer using our connection profile
	gw, err := gateway.Connect(
		gateway.WithConfig(cp),
		gateway.WithIdentity(gateway.NewX509Identity("Org1MSP", "/contracts/userid.pem")),
	)
	if err != nil {
		return nil, err
	}

	// Get network instance of our channel
	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		return nil, err
	}

	// Get instance of our chaincode from network
	contract := network.GetContract("nftcc")

	return &NFTService{
		Contract: contract,
	}, nil
}

func (service *NFTService) CreateNFT(name string, unique string, owner string) (string, error) {
	result, err := service.Contract.SubmitTransaction("CreateNFT", name, unique, owner)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func (service *NFTService) QueryNFT(name string) (*model.NFT, error) {
	result, err := service.Contract.EvaluateTransaction("QueryNFT", name)
	if err != nil {
		return nil, err
	}

	var nft model.NFT
	_ = json.Unmarshal(result, &nft)

	return &nft, nil
}
