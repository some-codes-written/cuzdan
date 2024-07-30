package common

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateToken() string {
	return uuid.New().String()
}

func ApiExceptions() []string {
	exceptions := []string{
		"auth",
	}

	return exceptions
}

func Extensions() []string {
	extensions := []string{
		".js",
		".css",
		".png",
		".jpg",
		".jpeg",
		".gif",
		".ico",
		".svg",
		".woff",
		".woff2",
		".ttf",
		".eot",
		".otf",
		".map",
		".json",
		".html",
		".txt",
	}
	return extensions
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if strings.Contains(a, e) {
			return true
		}
	}
	return false
}
