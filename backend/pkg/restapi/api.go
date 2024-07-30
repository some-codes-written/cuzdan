package restapi

import (
	"fmt"
	"immortality/pkg/domain/users"
	"immortality/pkg/restapi/controllers/models"
	"immortality/pkg/restapi/middleware"
	"immortality/pkg/restapi/routes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	PORT    = "8080"
	HOST    = "localhost"
	VERSION = "v1"
)

func Initialize() {
	fmt.Println("Initializing restapi...")

	r := mux.NewRouter()

	token, _ := dummyAutoLogin()
	amw := middleware.AuthMiddleWare{}
	amw.Init(token)
	r.Use(amw.Middleware)

	routes.Routes(r, VERSION)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%s", HOST, PORT),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Restapi initialized on http://localhost:" + PORT + "")

	log.Fatal(srv.ListenAndServe())
}

// geliştirme aşamasında kullanılacak
func dummyAutoLogin() (string, error) {
	authRequest := models.AuthRequest{
		Email:    "oguzhan.saricam@gmail.com",
		Password: "123456",
	}
	store := users.NewUserStore()
	_, err := store.VerifyCredential(authRequest.Email, authRequest.Password)
	if err != nil {
		return "", err
	}
	user, err := store.GetUserByEmail(authRequest.Email)
	if err != nil {
		return "", err
	}
	_, err = store.ExpireAllTokens()
	if err != nil {
		return "", err
	}
	token, err := store.GenerateToken(user)
	if err != nil {
		return "", err
	}
	return token.Token, nil
}
