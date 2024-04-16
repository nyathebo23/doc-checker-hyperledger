package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Create adds a new id with value to the world state
func (rc *DocumentSavesContract) Create(
	ctx contractapi.TransactionContextInterface
) error {

    // Get new asset from transient map
    transientMap, err := ctx.GetStub().GetTransient()
    if err != nil {
        return fmt.Errorf("error getting transient: %v", err)
    }

    // Asset properties are private, therefore they get passed in transient field, instead of func args
    transientDocJSON, ok := transientMap["doc_properties"]
    if !ok {
        //log error to stdout
        return fmt.Errorf("asset not found in the transient map input")
    }

	type DocumentSaveTransientInput struct {
		ID          string `json:"id"`
		OrganizationId   string    `json:"organization_id"`
		BeneficiaryId    string    `json:"beneficiary_id"`
		Filename    string `json:"filename"`
		Hash      	bool   `json:"hash"`
	}

    var docSaveInput DocumentSaveTransientInput
    err = json.Unmarshal(transientDocJSON, &docSaveInput)
    if err != nil {
        return fmt.Errorf("failed to unmarshal JSON: %v", err)
    }

    if len(docSaveInput.OrganizationId) == 0 {
        return fmt.Errorf("OrganizationId field must be a non-empty string")
    }
    if len(docSaveInput.ID) == 0 {
        return fmt.Errorf("assetID field must be a non-empty string")
    }
    if len(docSaveInput.BeneficiaryId) == 0 {
        return fmt.Errorf("BeneficiaryId field must be a non-empty string")
    }
    if len(docSaveInput.Filename) == 0 {
        return fmt.Errorf("Filename field must be a non-empty string")
    }
	if len(docSaveInput.Hash) == 0 {
        return fmt.Errorf("Hash field must be a non-empty string")
    }


	chainCodeOrgArgs := util.ToChaincodeArgs("Read", docSaveInput.OrganizationId)

	if res := ctx.GetStub().InvokeChaincode("organization", chainCodeArgs, ""); res.Status != 200 {
		return fmt.Errorf("Organization '%s' does not exist", docSaveInput.OrganizationId)
	}

    chainCodeBeneficiaryArgs := util.ToChaincodeArgs("Read", docSaveInput.BeneficiaryId)

	if res := ctx.GetStub().InvokeChaincode("beneficiary", chainCodeArgs, ""); res.Status != 200 {
		return fmt.Errorf("Beneficiary '%s' does not exist", docSaveInput.BeneficiaryId)
	}

	// Get collection name for this organization.
	orgCollection, err := getCollectionName(ctx)
	if err != nil {
		return fmt.Errorf("failed to infer private collection name for the org: %v", err)
	}
	
    // Check if asset already exists
    assetAsBytes, err := ctx.GetStub().GetPrivateData(orgCollection, id)
    if err != nil {
        return fmt.Errorf("failed to get asset: %v", err)
    } else if assetAsBytes != nil {
        fmt.Println("Asset already exists: " + id)
        return fmt.Errorf("this asset already exists: " + id)
    }

	existing, err := ctx.GetStub().GetState(id)

	if err != nil {
		return fmt.Errorf("Unable to interact with world state")
	}

	if existing != nil {
		return fmt.Errorf("Cannot create world state pair with id %s. Already exists", id)
	}

	newDocumentSave := &DocumentSave{
		ID:     id,
		Filename:      filename
		OrganizationId:   organizationId, // TODO: Verify this name is unique
		BeneficiaryId:   beneficiaryId,
		Hash: 	hash,
	}

    // Get ID of submitting client identity
    clientID, err := submittingClientIdentity(ctx)
    if err != nil {
        return err
    }

	newDocumentSaveasBytes, err := json.Marshal(newDocumentSave)
    if err != nil {
        return fmt.Errorf("failed to marshal asset into JSON: %v", err)
    }

    // Verify that the client is submitting request to peer in their organization
    // This is to ensure that a client from another org doesn't attempt to read or
    // write private data from this peer.
    err = verifyClientOrgMatchesPeerOrg(ctx)
    if err != nil {
        return fmt.Errorf("CreateAsset cannot be performed: Error %v", err)
    }

	// Put asset appraised value into owners org specific private data collection
	log.Printf("Put: collection %v, ID %v", orgCollection, id)
	err = ctx.GetStub().PutPrivateData(orgCollection, id, newDocumentSaveasBytes)
	if err != nil {
		return fmt.Errorf("failed to put asset private details: %v", err)
	}

	return nil
}
