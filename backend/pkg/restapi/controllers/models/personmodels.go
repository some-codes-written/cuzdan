package models

import "time"

type PersonDto struct {
	ID         uint      `json:"id"`
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

type PersonResponse struct {
	ApiResponse
	Person PersonDto `json:"person"`
}

type PersonUpdateRequest struct {
	Person PersonDto `json:"person"`
}

type PersonUpdateResponse struct {
	ApiResponse
	Person PersonDto `json:"person"`
}

type PersonDeleteRequest struct {
	ID uint `json:"id"`
}

type PersonDeleteResponse struct {
	ApiResponse
}

type PersonListResponse struct {
	ApiResponse
	Persons []PersonDto `json:"persons"`
}

type PersonCreateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Gsm       string `json:"gsm"`
	Gender    string `json:"gender"`
}

type PersonCreateResponse struct {
	ApiResponse

	Person PersonDto `json:"person"`
}
