package beneficiaries

import (
	"feuli/backend/models"
)

func Show(clients *hyperledger.Clients, id string) (beneficiary *models.Beneficiary, err error) {

	beneficiary := new(models.Beneficiary)

	res, err := clients.Query("enspy", "mainchannel", "beneficiary", "queryString", [][]byte{
		[]byte("{\"selector\":{ \"id\": { \"$eq\":\"" + id + "\" } }}"),
	})
	if err != nil {
		return
	}

	if err = json.Unmarshal(res, beneficiary); err != nil {
		return
	}

	list := *beneficiary

	if len(list) == 0 {
		err = errors.New("unable to find beneficiary with id " + id)
		return
	}

	beneficiary = &list[0]

	return
}
