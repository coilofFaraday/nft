package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// TransferOwnership transfers the ownership of an NFT
func (s *SmartContract) TransferOwnership(ctx contractapi.TransactionContextInterface, nft string, newOwner string) error {
	nftBytes, _ := ctx.GetStub().GetState(nft)

	var nftAsset NFT
	_ = json.Unmarshal(nftBytes, &nftAsset)

	nftAsset.Owner = newOwner

	nftBytes, _ = json.Marshal(nftAsset)

	return ctx.GetStub().PutState(nft, nftBytes)
}
