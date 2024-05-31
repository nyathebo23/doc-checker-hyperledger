package beneficiaries

import "feuli/backend/models"

func Store(clients *hyperledger.Clients, firstName string, lastName string, email string, password string) (beneficiary *models.Beneficiary, err error) {
	beneficiary, err = models.NewBeneficiary(firstName, lastName, email, password)
	if err != nil {
		return
	}

	packet, err := json.Marshal(beneficiary)
	if err != nil {
		return
	}

	if _, err = clients.Invoke("enspy", "beneficiary", "store", [][]byte{
		packet,
	}); err != nil {
		return
	}

	return
}
