package models

type Beneficiaries []Beneficiary

type Beneficiary struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func NewBeneficiary(firstName string, lastName string, email string, password string) (beneficiary *Beneficiary, err error) {
	beneficiary = new(Beneficiary)

	id, err := genUUID()
	if err != nil {
		return
	}

	beneficiary.ID = id
	beneficiary.FirstName = firstName
	beneficiary.LastName = lastName
	beneficiary.Email = email
	if beneficiary.Password, err = EncryptPassword(password); err != nil {
		return
	}

	return
}
