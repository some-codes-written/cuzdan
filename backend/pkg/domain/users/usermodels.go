package users

import (
	"immortality/pkg/common"
	"immortality/pkg/domain/persons"
	"time"
)

const (
	USERS             = "users"
	CREDENTIALS       = "credentials"
	CREDENTIALCHANGES = "credential_changes"
	USERTOKENS        = "user_tokens"
)

// User
type User struct {
	common.EntityModel
	PersonId      uint            `json:"personId" example:"1"`              // Person
	Person        *persons.Person `json:"person" gorm:"foreignKey:PersonId"` // Person
	Email         string          `json:"email" example:"john.doe@gmail.com"`
	Gsm           string          `json:"gsm" example:"555-555-5555"`
	FirstName     string          `json:"firstName" example:"John"`
	LastName      string          `json:"lastName" example:"Doe"`
	LastLoginDate time.Time       `json:"lastLoginDate" example:"2021-01-01T00:00:00Z"`
	UserToken     *[]UserToken    `json:"userToken" gorm:"foreignKey:UserId"`
	Credential    *[]Credential   `json:"credential" gorm:"foreignKey:UserId"`
}

// Credential
type Credential struct {
	common.EntityModel
	UserId   uint   `json:"userId" example:"1"`
	User     *User  `json:"user" gorm:"foreignKey:UserId"`
	Username string `json:"username" example:"ouzsrcm"`
	Email    string `json:"email" example:"john.doe@gmail.com"`
	Password string `json:"password" example:"123456"`
}

// CredentialChange
type CredentialChange struct {
	common.EntityModel
	UserId       uint        `json:"userId" example:"1"`
	User         *User       `json:"user" gorm:"foreignKey:UserId"` // User
	CredentialId string      `json:"credentialId" example:"1"`
	Credential   *Credential `json:"credential" gorm:"foreignKey:CredentialId"`
	Hash         string      `json:"hash" example:"UUID"`
	Url          string      `json:"url" example:"1"`
	ProcessDate  time.Time   `json:"processDate" example:"2021-01-01T00:00:00Z"`
	OldPassword  string      `json:"oldPassword" example:"123456"`
	NewPassword  string      `json:"newPassword" example:"1234567"`
	IsChanged    bool        `json:"isChanged" example:"true"`
}

// UserToken
type UserToken struct {
	common.EntityModel
	UserId         uint      `json:"userId" example:"1"`
	User           *User     `json:"user" gorm:"foreignKey:UserId"`
	Token          string    `json:"token" example:"UUID"` // UUID
	ExpirationDate time.Time `json:"expirationDate" example:"2021-01-01T00:00:00Z"`
	IsActive       bool      `json:"isActive" example:"true"`
}
