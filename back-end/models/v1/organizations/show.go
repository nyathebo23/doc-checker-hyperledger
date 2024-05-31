package organizations

import (
	"feuli/backend/models"
)

func Show(clients *hyperledger.Clients, id string) (organization *models.Organization, err error) {

	organization := new(models.Organization)

	res, err := clients.Query("enspy", "mainchannel", "organization", "queryString", [][]byte{
		[]byte("{\"selector\":{ \"id\": { \"$eq\":\"" + id + "\" } }}"),
	})
	if err != nil {
		return
	}

	if err = json.Unmarshal(res, organization); err != nil {
		return
	}

	list := *organization

	if len(list) == 0 {
		err = errors.New("unable to find organization with id " + id)
		return
	}

	organization = &list[0]

	return
}