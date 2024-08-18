package person_operations

import (
	"immortality/pkg/domain/persons/person_dtos"
)

func (ops *PersonOperations) GetPerson(id uint) (*person_dtos.PersonDto, error) {
	res, err := ops.Store.GetPerson(id)
	if err != nil {
		return nil, err
	}
	temp := person_dtos.PersonDto{
		FirstName:  res.FirstName,
		LastName:   res.LastName,
		Email:      res.Email,
		Phone:      res.Phone,
		Gsm:        res.Gsm,
		Gender:     res.Gender,
		BirthDate:  res.BirthDate,
		BirthPlace: res.BirthPlace,
		IdNumber:   res.IdNumber,
	}
	return &temp, nil
}

func (ops *PersonOperations) GetPersons() (*[]person_dtos.PersonDto, error) {
	res, err := ops.Store.GetPersons()
	if err != nil {
		return nil, err
	}
	temp := []person_dtos.PersonDto{}

	for _, item := range *res {
		temp = append(temp, person_dtos.PersonDto{
			FirstName:  item.FirstName,
			LastName:   item.LastName,
			Email:      item.Email,
			Phone:      item.Phone,
			Gsm:        item.Gsm,
			Gender:     item.Gender,
			BirthDate:  item.BirthDate,
			BirthPlace: item.BirthPlace,
			IdNumber:   item.IdNumber,
		})
	}

	return &temp, nil
}
