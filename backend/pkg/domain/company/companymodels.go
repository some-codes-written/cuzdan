package company

import (
	"immortality/pkg/common"
	"immortality/pkg/domain/persons"
	"time"
)

const (
	COMPANIES            = "companies"
	COMPANY_TYPES        = "company_types"
	COMPANY_PERSONS      = "company_people"
	COMPANY_PERSON_TYPES = "company_person_types"
)

type CompanyType struct {
	common.EntityModel
	Name        string `json:"name" example:"Company Type Name"`               // Company type name
	Description string `json:"description" example:"Company Type Description"` // Company type description
}

type Company struct {
	common.EntityModel

	CompanyTypeID uint            `json:"company_type_id"` // FK to CompanyType
	Name          string          `json:"name"`            // Company name
	Description   string          `json:"description"`     // Company description
	Email         string          `json:"email"`           // Company email
	Phone         string          `json:"phone"`           // Company phone
	Website       string          `json:"website"`         // Company website
	Address       string          `json:"address"`         // Company address
	IsActive      bool            `json:"is_active"`       // Is active
	CompanyType   *CompanyType    `json:"company_type"`    // FK to CompanyType
	AuthPersonId  uint            `json:"auth_person_id"`  // FK to Person
	AuthPerson    *persons.Person `json:"auth_person"`     // FK to Person

}

type CompanyPerson struct {
	common.EntityModel

	CompanyId  uint            `json:"company_id"`  // FK to Company
	PersonId   uint            `json:"person_id"`   // FK to User
	AssignDate time.Time       `json:"assign_date"` // Assign date
	Company    *Company        `json:"company"`     // FK to Company
	Person     *persons.Person `json:"person"`      // FK to User
}

type CompanyPersonType struct {
	common.EntityModel

	Name        string `json:"name" example:"Company Person Type Name"`               // Company person type name
	Description string `json:"description" example:"Company Person Type Description"` // Company person type description
}
