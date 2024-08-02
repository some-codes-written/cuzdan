package dtos

import (
	"immortality/pkg/common"
	"time"
)

type PersonDto struct {
	common.DTOModel

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

type AddressDto struct {
	common.DTOModel

	PersonId   int64      `json:"person_id"`
	Person     *PersonDto `json:"person"`
	CountryId  int64      `json:"country_id"`
	CityId     int64      `json:"city_id"`
	DistrictId int64      `json:"district_id"`
	Street     string     `json:"street"`
	ZipCode    string     `json:"zip_code"`
}
