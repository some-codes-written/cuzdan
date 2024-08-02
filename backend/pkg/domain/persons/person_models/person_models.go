package person_models

import (
	"immortality/pkg/common"
	"time"
)

const (
	PERSONS   = "people"
	ADDRESSES = "addresses"
)

// Person
type Person struct {
	common.EntityModel

	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Gsm        string    `json:"gsm"`
	Gender     string    `json:"gender"`
	BirthDate  time.Time `json:"birth_date"`
	BirthPlace string    `json:"birth_place"`
	IdNumber   string    `json:"id_number"`
}

// Address
type Address struct {
	common.EntityModel

	PersonId   int64   `json:"person_id"`
	Person     *Person `json:"person"`
	CountryId  int64   `json:"country_id"`
	CityId     int64   `json:"city_id"`
	DistrictId int64   `json:"district_id"`
	Street     string  `json:"street"`
	ZipCode    string  `json:"zip_code"`
}
