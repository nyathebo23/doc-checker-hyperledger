package organizations

import "feuli/backend/models"

func Index(clients *hyperledger.Clients) (organizations *models.Organizations, err error) {
	organizations = new(models.Organizations)

	res, err := clients.Query("enspy", "mainchannel", "organization", "queryString", [][]byte{
		[]byte(),
	})
	if err != nil {
		return
	}

	if err = json.Unmarshal(res, organizations); err != nil {
		return
	}

	return
}
