package restapi

import (
	"immortality/pkg/restapi/apibase"
	"immortality/pkg/restapi/apioperations/user_operations"
	"net/http"
)

// Index godoc
// @Summary Get all persons
// @Description Get all persons
// @Tags persons
// @Accept  json
// @Produce  json
// @Success 200 {object} models.PersonListResponse
// @Failure 400 {object} models.PersonListResponse
// @Failure 500 {object} models.PersonListResponse
// @Router /persons/list [get]
func PersonList(w http.ResponseWriter, r *http.Request) {

	_, res := user_operations.PersonList(w, r)

	apibase.ApiResult(w, r, apibase.ResultInfo{
		StatusCode:  http.StatusOK,
		ContentType: "application/json",
		Data:        res,
	})

}

// Index godoc
// @Summary Get a person
// @Description Get a person
// @Tags person
// @Accept  json
// @Produce  json
// @Success 200 {object} models.PersonResponse
// @Failure 400 {object} models.PersonResponse
// @Failure 500 {object} models.PersonResponse
// @Router /persons/get/{id} [get]
func GetPerson(w http.ResponseWriter, r *http.Request) {

	_, res := user_operations.GetPerson(w, r)

	apibase.ApiResult(w, r, apibase.ResultInfo{
		StatusCode:  http.StatusOK,
		ContentType: "application/json",
		Data:        res,
	})
}
