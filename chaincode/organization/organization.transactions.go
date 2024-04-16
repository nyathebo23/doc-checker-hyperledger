package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Transactions get all the transactions of an id
func (rc *OrganizationsContract) Transactions(
	ctx contractapi.TransactionContextInterface,
	id string,
) ([]*OrganizationTransactionItem, error) {
	historyIface, err := ctx.GetStub().GetHistoryForKey(id)
	if err != nil {
		return nil, err
	}

	var rets []*OrganizationTransactionItem
	for historyIface.HasNext() {
		val, err := historyIface.Next()
		if err != nil {
			return nil, err
		}

		var org Organization
		if err = json.Unmarshal(val.Value, &res); err != nil {
			return nil, err
		}

		rets = append(rets, &OrganizationTransactionItem{
			TXID:      val.TxId,
			Timestamp: int64(val.Timestamp.GetNanos()),
			Organization:  org,
		})
	}

	return rets, nil
}
