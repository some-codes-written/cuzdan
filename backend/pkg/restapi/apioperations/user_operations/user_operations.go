package user_operations

import (
	"encoding/json"
	"immortality/pkg/domain/users"
	"immortality/pkg/restapi/apibase"
	"immortality/pkg/restapi/controllers/models"
	"immortality/util"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func UserList(w http.ResponseWriter, r *http.Request) (bool, models.UserListResponse) {

	userStore := users.NewUserStore()
	res, err := userStore.GetUsers()

	var users []models.UserDto
	var response models.UserListResponse

	if err != nil {
		response.Status = apibase.ApiStatusError
		response.ErrorMessage = err.Error()
		return false, response
	}
	for _, user := range res {
		users = append(users, models.UserDto{
			ID:        user.ID,
			Email:     user.Email,
			Gsm:       user.Gsm,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		})
	}
	response.Users = users
	response.Status = apibase.ApiStatusSuccess
	return true, response
}

func CreateUser(w http.ResponseWriter, r *http.Request) (bool, models.UserCreateResponse) {

	var user users.User
	var credential users.Credential
	var request models.UserCreateRequest
	var response models.UserCreateResponse

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		response.Status = apibase.ApiStatusError
		response.ErrorMessage = err.Error()
		return false, response
	}
	user.Email = request.Email
	user.Gsm = request.Gsm
	user.FirstName = request.FirstName
	user.LastName = request.LastName

	userStore := users.NewUserStore()
	res, err := userStore.CreateUser(&user)

	if err != nil {
		response.Status = apibase.ApiStatusError
		response.ErrorMessage = err.Error()
		return false, response
	}

	credential.UserID = res.ID
	credential.Email = request.Email
	credential.Password = request.Password
	credential.CreatedAt = time.Now()

	creres, err := userStore.CreateCredential(&credential)

	if err != nil {
		response.Status = apibase.ApiStatusError
		response.ErrorMessage = err.Error()
		return false, response
	}

	response.Credential = models.CredentialDto{
		ID:     creres.ID,
		UserId: creres.UserID,
		Email:  creres.Email,
		//Password:  creres.Password
	}
	response.User = models.UserDto{
		ID:        res.ID,
		Email:     res.Email,
		Gsm:       res.Gsm,
		FirstName: res.FirstName,
		LastName:  res.LastName,
	}
	response.Status = apibase.ApiStatusSuccess
	return true, response
}

func GetUser(w http.ResponseWriter, r *http.Request) (bool, models.UserResponse) {

	response := models.UserResponse{}
	strId := mux.Vars(r)["id"]
	id, err := util.ToUint(strId)
	if err != nil {
		response.Status = apibase.ApiStatusError
		response.ErrorMessage = err.Error()
		return false, response
	}
	userStore := users.NewUserStore()
	res, err := userStore.GetUser(id)
	if err != nil {
		response.Status = apibase.ApiStatusError
		response.ErrorMessage = err.Error()
		return false, response
	}
	response.User = models.UserDto{
		ID:        res.ID,
		Email:     res.Email,
		Gsm:       res.Gsm,
		FirstName: res.FirstName,
		LastName:  res.LastName,
	}
	response.Status = apibase.ApiStatusSuccess

	return true, response

}

func UpdateUser(w http.ResponseWriter, r *http.Request) (bool, models.UserUpdateResponse) {
	var user models.UserUpdateRequest
	response := models.UserUpdateResponse{}
	strId := r.URL.Query().Get("id")
	id, err := util.ToUint(strId)
	if err != nil {
		response.Status = apibase.ApiStatusError
		response.ErrorMessage = err.Error()
		return false, response
	}
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Status = apibase.ApiStatusError
		response.ErrorMessage = err.Error()
		return false, response
	}
	userStore := users.NewUserStore()
	res, _ := userStore.GetUser(id)

	res.Email = user.Email
	res.Gsm = user.Gsm
	res.FirstName = user.FirstName
	res.LastName = user.LastName

	res, err = userStore.UpdateUser(res)
	if err != nil {
		response.Status = apibase.ApiStatusError
		response.ErrorMessage = err.Error()
		return false, response
	}
	response.User = models.UserDto{
		ID:        res.ID,
		Email:     res.Email,
		Gsm:       res.Gsm,
		FirstName: res.FirstName,
		LastName:  res.LastName,
	}
	response.Status = apibase.ApiStatusSuccess
	return true, response
}

func DeleteUser(w http.ResponseWriter, r *http.Request) (bool, models.UserDeleteResponse) {
	response := models.UserDeleteResponse{}
	strId := r.URL.Query().Get("id")
	id, err := util.ToUint(strId)
	if err != nil {
		response.Status = apibase.ApiStatusError
		response.ErrorMessage = err.Error()
		return false, response
	}
	userStore := users.NewUserStore()
	err = userStore.DeleteUser(id)
	if err != nil {
		response.Status = apibase.ApiStatusError
		response.ErrorMessage = err.Error()
		return false, response
	}
	response.Status = apibase.ApiStatusSuccess
	return true, response
}
