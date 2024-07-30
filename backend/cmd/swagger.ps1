# Path: backend\cmd\swagger.ps1

rm -r -fo ./docs

swag init -g .\pkg\restapi\controllers\users.go .\pkg\restapi\controllers\auth.go