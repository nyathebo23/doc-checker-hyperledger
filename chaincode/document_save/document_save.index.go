package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Index - read all resources from the world state
func (rc *DocumentSavesContract) Index(
	ctx contractapi.TransactionContextInterface,
) (rets []*DocumentSave, err error) {
	mspID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return
	}

	resultsIterator, _, err := ctx.GetStub().GetQueryResultWithPagination(`{"selector": {"id":{"$ne":"-"},"owner":"`+mspID+`"}}`, 0, "")
	if err != nil {
		return
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		queryResponse, err2 := resultsIterator.Next()
		if err2 != nil {
			return nil, err2
		}

		docSave := new(DocumentSave)
		if err = json.Unmarshal(queryResponse.Value, docSave); err != nil {
			return
		}

		rets = append(rets, docSave)
	}

	return
}
