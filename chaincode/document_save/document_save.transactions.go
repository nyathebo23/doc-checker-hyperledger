package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Transactions get all the transactions of an id
func (rc *DocumentSavesContract) Transactions(
	ctx contractapi.TransactionContextInterface,
	id string,
) ([]*DocumentSaveTransactionItem, error) {
	historyIface, err := ctx.GetStub().GetHistoryForKey(id)
	if err != nil {
		return nil, err
	}

	var rets []*DocumentSaveTransactionItem
	for historyIface.HasNext() {
		val, err := historyIface.Next()
		if err != nil {
			return nil, err
		}

		var docSave DocumentSave
		if err = json.Unmarshal(val.Value, &res); err != nil {
			return nil, err
		}

		rets = append(rets, &DocumentSaveTransactionItem{
			TXID:      val.TxId,
			Timestamp: int64(val.Timestamp.GetNanos()),
			DocumentSave:  docSave,
		})
	}

	return rets, nil
}
