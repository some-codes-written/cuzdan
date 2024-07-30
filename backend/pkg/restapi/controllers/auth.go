// path: backend\pkg\restapi\restapi.go

//TODO: redis kullanılacak

package restapi

import (
	"encoding/json"
	"immortality/pkg/domain/users"
	"immortality/pkg/restapi/apibase"
	"immortality/pkg/restapi/controllers/models"
	"net/http"
)

// / Register godoc
// @Summary auth for token
// @Description auth for token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body models.AuthRequest true "request"
// @Success 200 {object} models.AuthResponse "true"
// @Failure 401 {object} models.AuthResponse "false"
// @Failure 404 {object} models.AuthResponse "false"
// @Failure 500 {object} models.AuthResponse "false"
// @Router /auth [post]
func Auth(w http.ResponseWriter, r *http.Request) {

	var response models.AuthResponse
	var model models.AuthRequest
	var token string

	json.NewDecoder(r.Body).Decode(&model)

	userStore := users.NewUserStore()

	res, err := userStore.VerifyCredential(model.Email, model.Password)
	if err != nil {
		response.Status = apibase.ApiStatusError
		response.ErrorMessage = err.Error()
		apibase.ApiResult(w, r, apibase.ResultInfo{
			StatusCode:  http.StatusNotFound,
			ContentType: "application/json",
			Data:        response,
		})
		return
	}
	user, err := userStore.GetUserByEmail(model.Email)
	if err != nil {
		response.Status = apibase.ApiStatusError
		response.ErrorMessage = err.Error()
		apibase.ApiResult(w, r, apibase.ResultInfo{
			StatusCode:  http.StatusNotFound,
			ContentType: "application/json",
			Data:        response,
		})
		return
	}
	if res {
		tokenres, err := userStore.GenerateToken(user)
		if err != nil {
			response.Status = apibase.ApiStatusError
		} else {
			response.Status = apibase.ApiStatusSuccess
		}
		token = tokenres.Token
	} else {
		response.Status = apibase.ApiStatusError
		response.ErrorMessage = "Invalid email or password"
		apibase.ApiResult(w, r, apibase.ResultInfo{
			StatusCode:  http.StatusBadRequest,
			ContentType: "application/json",
			Data:        response,
		})
		return
	}
	response.Data = models.AuthResponseData{
		Token:  token,
		UserId: user.ID,
	}
	apibase.ApiResult(w, r, apibase.ResultInfo{
		StatusCode:  http.StatusOK,
		ContentType: "application/json",
		Data:        response,
	})
}

// / ExpireToken godoc
// @Summary expire token
// @Description expire token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body models.ExpireTokenRequest true "request"
// @Success 200 {object} models.ExpireTokenResponse "true"
// @Failure 401 {object} models.ExpireTokenResponse "false"
// @Failure 404 {object} models.ExpireTokenResponse "false"
// @Failure 500 {object} models.ExpireTokenResponse "false"
// @Router /auth/expire_token [post]
func ExpireToken(w http.ResponseWriter, r *http.Request) {

	var response models.ExpireTokenResponse
	var model models.ExpireTokenRequest
	json.NewDecoder(r.Body).Decode(&model)

	userStore := users.NewUserStore()

	res, err := userStore.ExpireToken(model.Token)
	if err != nil {
		response.Status = apibase.ApiStatusError
		response.ErrorMessage = err.Error()
		apibase.ApiResult(w, r, apibase.ResultInfo{
			ContentType: "application/json",
			StatusCode:  http.StatusNotFound,
			Data:        response,
		})
		return
	}
	if res != nil {
		response.Status = apibase.ApiStatusSuccess
		response.ErrorMessage = "Token Expired"
	} else {
		response.Status = apibase.ApiStatusError
		response.ErrorMessage = "Token not found"
	}
	apibase.ApiResult(w, r, apibase.ResultInfo{
		ContentType: "application/json",
		StatusCode:  http.StatusOK,
		Data:        response,
	})
}

// / TokenExists godoc
// @Summary token exists
// @Description token exists
// @Tags auth
// @Accept json
// @Produce json
// @Param request body models.TokenExistsRequest true "request"
// @Success 200 {object} models.TokenExistsResponse "true"
// @Failure 401 {object} models.TokenExistsResponse "false"
// @Failure 404 {object} models.TokenExistsResponse "false"
// @Failure 500 {object} models.TokenExistsResponse "false"
// @Router /auth/token_exists [post]
func TokenExists(w http.ResponseWriter, r *http.Request) {
	var response models.TokenExistsResponse
	var model models.TokenExistsRequest
	json.NewDecoder(r.Body).Decode(&model)

	userStore := users.NewUserStore()

	res, err := userStore.TokenExists(model.Token)
	if err != nil {
		response.Status = apibase.ApiStatusError
		response.ErrorMessage = err.Error()
		resultInfo := apibase.NewResultInfo(http.StatusNotFound, err.Error(), "application/json", response)
		apibase.ApiResult(w, r, *resultInfo)
		return
	}
	response.Exists = res
	response.Status = apibase.ApiStatusSuccess
	apibase.ApiResult(w, r, apibase.ResultInfo{
		ContentType: "application/json",
		StatusCode:  http.StatusOK,
		Data:        response,
	})
}

// / CurrentTokens godoc
// @Summary current tokens
// @Description current tokens
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} models.CurrentTokensResponse "true"
// @Failure 401 {object} models.CurrentTokensResponse "false"
// @Failure 404 {object} models.CurrentTokensResponse "false"
// @Failure 500 {object} models.CurrentTokensResponse "false"
// @Router /auth/current_tokens [get]
func CurrentTokens(w http.ResponseWriter, r *http.Request) {
	var response models.CurrentTokensResponse
	userStore := users.NewUserStore()
	res, err := userStore.GetTokens()
	if err != nil {
		response.Status = apibase.ApiStatusError
		response.ErrorMessage = err.Error()
		apibase.ApiResult(w, r, apibase.ResultInfo{
			ContentType: "application/json",
			StatusCode:  http.StatusNotFound,
			Data:        response,
		})
		return
	}
	tokens := make(map[string]uint)
	for _, token := range res {
		tokens[token.Token] = token.UserId
	}
	response.Data = tokens
	response.Status = apibase.ApiStatusSuccess
	apibase.ApiResult(w, r, apibase.ResultInfo{
		ContentType: "application/json",
		StatusCode:  http.StatusOK,
		Data:        response,
	})
}

// / ExpireAllTokens godoc
// @Summary expire all tokens
// @Description expire all tokens
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} models.ExpireAllTokensResponse "true"
// @Failure 401 {object} models.ExpireAllTokensResponse "false"
// @Failure 404 {object} models.ExpireAllTokensResponse "false"
// @Failure 500 {object} models.ExpireAllTokensResponse "false"
// @Router /auth/expire_all_tokens [get]
func ExpireAllTokens(w http.ResponseWriter, r *http.Request) {
	userStore := users.NewUserStore()
	res, err := userStore.ExpireAllTokens()
	if err != nil {
		response := &models.ExpireAllTokensResponse{}
		response.Status = apibase.ApiStatusError
		apibase.ApiResult(w, r, apibase.ResultInfo{
			ContentType: "application/json",
			StatusCode:  http.StatusNotFound,
			Data:        response,
		})
		return
	}
	var resultInfo apibase.ResultInfo
	resultInfo.ContentType = "application/json"
	if res {
		resultInfo.StatusCode = http.StatusOK
	} else {
		resultInfo.StatusCode = http.StatusBadRequest
	}
	apibase.ApiResult(w, r, resultInfo)
}
