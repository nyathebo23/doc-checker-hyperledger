package main

// https://hyperledger-fabric.readthedocs.io/en/latest/chaincode4ade.html
import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	cc, err := contractapi.NewChaincode(&OrganizationsContract{})

	if err != nil {
		panic(err.Error())
	}

	if err := cc.Start(); err != nil {
		panic(err.Error())
	}
}

// ResourcesContract contract for handling writing and reading from the world state
type OrganizationsContract struct {
	contractapi.Contract
}

// InitLedger adds a base set of cars to the ledger
func (rc *OrganizationsContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	return nil
}
