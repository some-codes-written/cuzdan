package main

import (
	"immortality/pkg/initializer"
	"immortality/pkg/restapi"
)

func main() {

	initializer.Godoc("./swagger.ps1")

	initializer.Initialize()

	restapi.Initialize()

}
