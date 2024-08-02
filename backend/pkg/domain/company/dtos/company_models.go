package dtos

import "immortality/pkg/common"

type CompanyDto struct {
	common.DTOModel
	CompanyTypeID uint   `json:"company_type_id"` // FK to CompanyType
	Name          string `json:"name"`            // Company name
	Description   string `json:"description"`     // Company description
	Email         string `json:"email"`           // Company email
	Phone         string `json:"phone"`           // Company phone
	Website       string `json:"website"`         // Company website
	Address       string `json:"address"`         // Company address
	IsActive      bool   `json:"is_active"`       // Is active
	AuthPersonId  uint   `json:"auth_person_id"`  // FK to Person
}

type CompanyListDto struct {
	common.DTOModel
	Companies []CompanyDto `json:"companies"`
}
