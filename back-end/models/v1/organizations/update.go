package organizations

import (
	"errors"

	"feuli/backend/models"
)

type UpdateOpts struct {
	Replace bool
}

func Update(id string, org *models.Organization, opts *UpdateOpts) (*models.Organization, error) {

	if !opts.Replace {
		existingOrg, err := Show(clients, id)
		if err != nil {
			return nil, err
		}

		if org.Title != "" && existingOrg.Title != org.Title {
			existingOrg.Title = org.Title
		}

		if org.Name != "" && existingOrg.Name != org.Name {
			existingOrg.Name = org.Name
		}

		if org.IsActive != "" && existingOrg.IsActive != org.IsActive {
			existingOrg.IsActive = org.IsActive
		}

		if org.OrgType != "" && existingOrg.OrgType != org.OrgType {
			existingOrg.OrgType = org.OrgType
		}

	}

	packet, err := json.Marshal(org)
	if err != nil {
		return nil, err
	}

	if _, err = clients.Invoke("enspy", "organization", "update", [][]byte{
		[]byte(id),
		packet,
	}); err != nil {
		return nil, err
	}

}
