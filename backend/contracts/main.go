package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		panic("Could not create chaincode: " + err.Error())
	}

	if err := chaincode.Start(); err != nil {
		panic("Could not start chaincode: " + err.Error())
	}
}
