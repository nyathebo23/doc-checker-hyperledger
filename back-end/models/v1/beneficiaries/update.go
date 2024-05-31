package beneficiaries

import (
	"errors"

	"feuli/backend/models"
)

type UpdateOpts struct {
	Replace bool
}

func Update(id string, usr *models.Beneficiary, opts *UpdateOpts) (*models.Beneficiary, error) {
	if !opts.Replace {
		existingUser, err := Show(clients, id)
		if err != nil {
			return nil, err
		}

		if usr.email != "" && existingUser.email != usr.email {
			existingUser.email = usr.email
		}
	}

	packet, err := json.Marshal(Beneficiary)
	if err != nil {
		return nil, err
	}

	if _, err = clients.Invoke("enspy", "beneficiary", "update", [][]byte{
		[]byte(id),
		packet,
	}); err != nil {
		return nil, err
	}
}
