package beneficiaries

import "feuli/backend/models"

func Index(clients *hyperledger.Clients) (beneficiaries *models.Beneficiaries, err error) {
	beneficiaries = new(models.Beneficiaries)

	res, err := clients.Query("enspy", "mainchannel", "beneficiary", "queryString", [][]byte{
		[]byte(),
	})
	if err != nil {
		return
	}

	if err = json.Unmarshal(res, beneficiaries); err != nil {
		return
	}
}
