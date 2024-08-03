package company_dtos

import (
	"immortality/pkg/common"
	"time"
)

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

type CompanyTypeDto struct {
	common.DTOModel
	Name        string `json:"name" example:"Company Type Name"`               // Company type name
	Description string `json:"description" example:"Company Type Description"` // Company type description
}

type CompanyTypeListDto struct {
	common.DTOModel
	CompanyTypes []CompanyTypeDto `json:"company_types"`
}

type CompanyPersonDto struct {
	common.DTOModel

	CompanyId  uint      `json:"company_id"`  // FK to Company
	PersonId   uint      `json:"person_id"`   // FK to User
	AssignDate time.Time `json:"assign_date"` // Assign date
}
