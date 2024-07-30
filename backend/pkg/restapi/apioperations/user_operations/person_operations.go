package user_operations

import (
	"immortality/pkg/domain/persons"
	"immortality/pkg/restapi/apibase"
	"immortality/pkg/restapi/controllers/models"
	"immortality/util"
	"net/http"

	"github.com/gorilla/mux"
)

func PersonList(w http.ResponseWriter, r *http.Request) (bool, models.PersonListResponse) {

	personStore := persons.NewPersonStore()
	res, err := personStore.GetPersons()

	var persons []models.PersonDto
	var response models.PersonListResponse

	if err != nil {
		response.Status = apibase.ApiStatusError
		response.ErrorMessage = err.Error()
		return false, response
	}
	for _, person := range res {
		persons = append(persons, models.PersonDto{
			ID:        person.ID,
			FirstName: person.FirstName,
			LastName:  person.LastName,
		})
	}
	response.Persons = persons
	response.Status = apibase.ApiStatusSuccess
	return true, response

}

func GetPerson(w http.ResponseWriter, r *http.Request) (bool, models.PersonResponse) {

	response := models.PersonResponse{}
	strId := mux.Vars(r)["id"]
	id, err := util.ToUint(strId)
	if err != nil {
		response.Status = apibase.ApiStatusError
		response.ErrorMessage = err.Error()
		return false, response
	}
	personStore := persons.NewPersonStore()
	res, err := personStore.GetPerson(id)
	if err != nil {
		response.Status = apibase.ApiStatusError
		response.ErrorMessage = err.Error()
		return false, response
	}
	response.Person = models.PersonDto{
		ID:        res.ID,
		Email:     res.Email,
		Gsm:       res.Gsm,
		FirstName: res.FirstName,
		LastName:  res.LastName,
	}
	response.Status = apibase.ApiStatusSuccess

	return true, response

}
