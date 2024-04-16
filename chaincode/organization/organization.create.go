package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Create adds a new id with value to the world state
func (rc *OrganizationContract) Create(
	ctx contractapi.TransactionContextInterface,
	id string,
	title string,
	name string,
	active bool,
) error {
	if id == "0" || len(id) == 0 {
		id = uuid.New().String()
	}

	existing, err := ctx.GetStub().GetState(id)

	if err != nil {
		return fmt.Errorf("Unable to interact with world state")
	}

	if existing != nil {
		return fmt.Errorf("Cannot create world state pair with id %s. Already exists", id)
	}

	newOrganization := &Organization{
		ID:     id,
		Name:   name
		Title:   title, // TODO: Verify this name is unique
		IsActive:   active
	}

	bytes, err := json.Marshal(newOrganization)
	if err != nil {
		return fmt.Errorf("Unable to marshal object")
	}

	if err = ctx.GetStub().PutState(id, bytes); err != nil {
		return fmt.Errorf("Unable to interact with world state")
	}

	return nil
}
