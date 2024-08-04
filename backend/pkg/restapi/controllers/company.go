package restapi

import (
	"immortality/pkg/domain/company/company_operations"
	"immortality/pkg/restapi/apibase"
	"net/http"
)

func CompanyList(w http.ResponseWriter, r *http.Request) {

	_, res := company_operations.NewCompanyOperations().GetCompanyList()

	apibase.ApiResult(w, r, apibase.ResultInfo{
		StatusCode:  http.StatusOK,
		ContentType: "application/json",
		Data:        res,
	})
}
