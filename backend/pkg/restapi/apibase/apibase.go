// path: backend\pkg\restapi\apibase\apibase.go

package apibase

import (
	"encoding/json"
	"net/http"
)

const (
	ApiStatusSuccess = "success"
	ApiStatusError   = "error"
)

type ResultInfo struct {
	StatusCode  int         `json:"status_code"`  // HTTP status code
	StatusText  string      `json:"status_text"`  // HTTP status text
	Message     string      `json:"message"`      // Hata mesajı
	Data        interface{} `json:"data"`         // Veri
	ContentType string      `json:"content_type"` // Content-Type
}

// / ResultInfo nesnesi oluşturur
// / @param code HTTP status code
// / @param message Hata mesajı
// / @param content_type Content-Type
// / @param data Veri
func NewResultInfo(code int, message string, content_type string, data interface{}) *ResultInfo {

	resultInfo := new(ResultInfo)

	resultInfo.StatusCode = code
	resultInfo.StatusText = http.StatusText(code)
	resultInfo.Message = message
	resultInfo.Data = data
	resultInfo.ContentType = content_type

	return resultInfo

}

// / HTTP status kodunu ayarlar
// / @param code HTTP status code
func (s *ResultInfo) SetStatusText(code int) {
	s.StatusText = http.StatusText(code)
}

// / Response iletmek için kullanılır
// / @param w http.ResponseWriter
// / @param r *http.Request
// / @param i ResultInfo
func ApiResult(w http.ResponseWriter, r *http.Request, i ResultInfo) (bool, error) {

	if i.StatusCode >= 400 {
		return ApiResultError(w, r, i)
	}

	return ApiResultSuccess(w, r, i)

}

// / Hatalı response iletmek için kullanılır
// / @param w http.ResponseWriter
// / @param r *http.Request
// / @param i ResultInfo
func ApiResultError(w http.ResponseWriter, _ *http.Request, i ResultInfo) (bool, error) {

	w.Header().Set("Content-Type", i.ContentType)
	w.WriteHeader(i.StatusCode)

	json, _ := json.Marshal(i.Data)
	w.Write([]byte(json))

	return true, nil
}

// / Başarılı response iletmek için kullanılır
// / @param w http.ResponseWriter
// / @param r *http.Request
// / @param i ResultInfo
func ApiResultSuccess(w http.ResponseWriter, _ *http.Request, i ResultInfo) (bool, error) {

	w.Header().Set("Content-Type", i.ContentType)
	w.WriteHeader(i.StatusCode)

	json, _ := json.Marshal(i.Data)
	w.Write([]byte(json))

	return true, nil
}
