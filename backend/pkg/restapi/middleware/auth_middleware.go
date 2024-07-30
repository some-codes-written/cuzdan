package middleware

import (
	"fmt"
	"immortality/pkg/common"
	"immortality/pkg/domain/users"
	"log"
	"net/http"
	"strings"
)

type AuthMiddleWare struct {
	tokens map[string]uint
}

var (
	CURRENT_TOKEN string
)

func (m *AuthMiddleWare) Init(token string) {
	CURRENT_TOKEN = token
	log.Println("AuthMiddleWare init")
	log.Println("CURRENT_TOKEN: " + CURRENT_TOKEN)
}

func (m *AuthMiddleWare) Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if common.Contains(common.Extensions(), r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}
		if strings.Contains(r.URL.Path, "auth") {
			next.ServeHTTP(w, r)
			return
		}
		if strings.Contains(r.URL.Path, "swagger") {
			next.ServeHTTP(w, r)
			return
		}
		var token string
		if CURRENT_TOKEN != "" /*&& r.Header.Get("Authorization") == CURRENT_TOKEN*/ {
			token = CURRENT_TOKEN
			fmt.Println("Authorization Type: CURRENT_TOKEN")
		} else {
			token = r.Header.Get("Authorization")
		}
		log.Println("Request", r.URL.Path)
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Println("token: " + token)

		userStore := users.NewUserStore()
		tokenExists, _ := userStore.TokenExists(token)
		if !tokenExists {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
