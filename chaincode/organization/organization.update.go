package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Update changes the value with id in the world state
func (rc *OrganizationsContract) Update(ctx contractapi.TransactionContextInterface, id string, name string, title string, active bool) error {
	existing, err := ctx.GetStub().GetState(id)

	if err != nil {
		return fmt.Errorf("Unable to interact with world state")
	}

	if existing == nil {
		return fmt.Errorf("Cannot update world state pair with id %s. Does not exist", id)
	}

	var existingOrganization *Organization
	if err = json.Unmarshal(existing, &existingOrganization); err != nil {
		return fmt.Errorf("Unable to unmarshal existing into object")
	}
	existingOrganization.Name = name
	existingOrganization.Title = title		
	existingOrganization.IsActive = active

	newValue, err := json.Marshal(existingOrganization)
	if err != nil {
		return fmt.Errorf("Unable to marshal new object")
	}

	if err = ctx.GetStub().PutState(id, newValue); err != nil {
		return fmt.Errorf("Unable to interact with world state")
	}

	return nil
}
