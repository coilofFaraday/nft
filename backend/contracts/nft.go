package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// NFT defines the structure for an NFT
type NFT struct {
	Name   string `json:"name"`
	Unique string `json:"unique"`
	Owner  string `json:"owner"`
}

// CreateNFT creates a new NFT on the ledger
func (s *SmartContract) CreateNFT(ctx contractapi.TransactionContextInterface, name string, unique string, owner string) error {
	nft := NFT{
		Name:   name,
		Unique: unique,
		Owner:  owner,
	}

	nftBytes, _ := json.Marshal(nft)

	return ctx.GetStub().PutState(name, nftBytes)
}

// QueryNFT returns the NFT stored in the world state with given id
func (s *SmartContract) QueryNFT(ctx contractapi.TransactionContextInterface, name string) (*NFT, error) {
	nftBytes, err := ctx.GetStub().GetState(name)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state: %v", err)
	}

	var nft NFT
	_ = json.Unmarshal(nftBytes, &nft)

	return &nft, nil
}
