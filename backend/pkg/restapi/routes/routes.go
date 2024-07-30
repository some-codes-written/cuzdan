package routes

import (
	"fmt"

	"immortality/docs"
	restapi "immortality/pkg/restapi/controllers"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router, v string) {

	var pathbase string = fmt.Sprintf("/api/%s/", v)

	///swagger
	docs.SwaggerInfo.Title = "Immortality API"
	docs.SwaggerInfo.Description = "This is a Immortality server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.PathPrefix("/api/v1/swagger").Handler(httpSwagger.WrapHandler)

	///api/v1
	usersRoute := r.PathPrefix(pathbase + "users").Subrouter()
	usersRoute.HandleFunc("/list", restapi.UserList).Methods("GET")
	usersRoute.HandleFunc("/create", restapi.Create).Methods("POST")
	usersRoute.HandleFunc("/delete", restapi.Delete).Methods("DELETE")
	usersRoute.HandleFunc("/get/{id}", restapi.Get).Methods("GET")
	usersRoute.HandleFunc("/update/{id:[0-9]+}", restapi.Update).Methods("PUT")

	authRoute := r.PathPrefix(pathbase + "auth").Subrouter()
	authRoute.HandleFunc("", restapi.Auth).Methods("POST")
	authRoute.HandleFunc("/expire_token", restapi.ExpireToken).Methods("POST")
	authRoute.HandleFunc("/token_exists", restapi.TokenExists).Methods("POST")
	authRoute.HandleFunc("/current_tokens", restapi.CurrentTokens).Methods("GET")
	authRoute.HandleFunc("/expire_all_tokens", restapi.ExpireAllTokens).Methods("GET")

	personRoute := r.PathPrefix(pathbase + "persons").Subrouter()
	personRoute.HandleFunc("/list", restapi.PersonList).Methods("GET")
	personRoute.HandleFunc("/get/{id:[0-9]+}", restapi.GetPerson).Methods("GET")

}
