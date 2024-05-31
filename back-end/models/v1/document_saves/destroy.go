package document_saves

import (
	"feuli/backend/hyperledger"
)

func Destroy(clients *hyperledger.Clients, id string) error {
	res, err := Show(clients, id)
	if err != nil {
		return err
	}

	//res.Visible = false

	//if _, err = Update(clients, id, res, &UpdateOpts{ Replace: true }); err != nil {
		//return err
	//}

	return nil
}
