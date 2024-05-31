package organizations

import "feuli/backend/models"

func Store(clients *hyperledger.Clients, title string, name string, orgType string) (org *models.Organization, err error) {
	org, err = models.NewOrganization(title, name, orgType)
	if err != nil {
		return
	}

	packet, err := json.Marshal(org)
	if err != nil {
		return
	}

	if _, err = clients.Invoke("enspy", "organization", "store", [][]byte{
		packet,
	}); err != nil {
		return
	}

	return
}
