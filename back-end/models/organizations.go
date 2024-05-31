package models

import "time"

type Organizations []Organization

type Organization struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
	OrgType  string `json:"org_type"`
}


func NewOrganization(title string, name string, orgType string) (organization *Organization, err error) {
	organization = new(Organization)

	if organization.ID, err = genUUID(); err != nil {
		return
	}

	organization.Title = title
	organization.Name = name
	organization.OrgType = orgType
	organization.IsActive = false

	return
}
