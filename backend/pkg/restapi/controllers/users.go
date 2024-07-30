package restapi

import (
	"immortality/pkg/restapi/apibase"
	"immortality/pkg/restapi/apioperations/user_operations"
	"net/http"
)

// TODO: apibase.NewResultInfo metodu hatalı.
/// düzeltilmesi gerekiyor.

// Index godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.UserListResponse
// @Failure 400 {object} models.UserListResponse
// @Failure 500 {object} models.UserListResponse
// @Router /users/list [get]
func UserList(w http.ResponseWriter, r *http.Request) {

	_, res := user_operations.UserList(w, r)

	apibase.ApiResult(w, r, apibase.ResultInfo{
		StatusCode:  http.StatusOK,
		ContentType: "application/json",
		Data:        res,
	})

}

// Create godoc
// @Summary Create a user
// @Description Create a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.UserCreateRequest true "User"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.UserResponse
// @Failure 500 {object} models.UserResponse
// @Router /users/create [post]
func Create(w http.ResponseWriter, r *http.Request) {

	_, res := user_operations.CreateUser(w, r)

	apibase.ApiResult(w, r, apibase.ResultInfo{
		StatusCode:  http.StatusOK,
		ContentType: "application/json",
		Data:        res,
	})

}

// Get godoc
// @Summary Get a user
// @Description Get a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path uint true "User ID"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.UserResponse
// @Failure 500 {object} models.UserResponse
// @Router /users/get/{id} [get]
func Get(w http.ResponseWriter, r *http.Request) {

	_, res := user_operations.GetUser(w, r)

	apibase.ApiResult(w, r, apibase.ResultInfo{
		StatusCode:  http.StatusBadRequest,
		ContentType: "application/json",
		Data:        res,
	})

}

// Update godoc
// @Summary Update a user
// @Description Update a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path uint true "User ID"
// @Param user body models.UserUpdateRequest true "User"
// @Success 200 {object} models.UserUpdateResponse
// @Failure 400 {object} models.UserUpdateResponse
// @Failure 500 {object} models.UserUpdateResponse
// @Router /users/update/{id} [put]
func Update(w http.ResponseWriter, r *http.Request) {

	_, res := user_operations.UpdateUser(w, r)

	apibase.ApiResult(w, r, apibase.ResultInfo{
		StatusCode:  http.StatusBadRequest,
		ContentType: "application/json",
		Data:        res,
	})

}

// Delete godoc
// @Summary Delete a user
// @Description Delete a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path uint true "User ID"
// @Success 200 {object} models.UserDeleteResponse
// @Failure 400 {object} models.UserDeleteResponse
// @Failure 500 {object} models.UserDeleteResponse
// @Router /users/delete/{id} [delete]
func Delete(w http.ResponseWriter, r *http.Request) {

	_, res := user_operations.DeleteUser(w, r)

	apibase.ApiResult(w, r, apibase.ResultInfo{
		StatusCode:  http.StatusBadRequest,
		ContentType: "application/json",
		Data:        res,
	})
}
