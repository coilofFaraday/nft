package service

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type OwnershipService struct {
	Contract *gateway.Contract
}

func NewOwnershipService() (*OwnershipService, error) {
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

	return &OwnershipService{
		Contract: contract,
	}, nil
}

func (service *OwnershipService) TransferOwnership(nft string, newOwner string) (string, error) {
	result, err := service.Contract.SubmitTransaction("TransferOwnership", nft, newOwner)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
