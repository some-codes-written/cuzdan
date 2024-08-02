package users

import (
	"immortality/database"
	"immortality/pkg/domain/users/user_models"
)

func SetupDatabase() {

	dst := GetInterfaces()

	MigrateDatabase(dst...)

	SeedingDatabase()
}

func SeedingDatabase() {

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	SeedUser(db)

}

func GetInterfaces() []interface{} {
	models := make([]interface{}, 0)
	return append(models,
		&user_models.User{},
		&user_models.Credential{},
		&user_models.CredentialChange{},
		&user_models.UserToken{},
	)
}

func MigrateDatabase(dst ...interface{}) error {
	if dst != nil {
		dst = append(dst, GetInterfaces()...)
	}
	res := database.Migrate(dst...)
	return res
}
