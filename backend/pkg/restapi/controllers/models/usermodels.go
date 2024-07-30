package models

// / Models for user.go
type UserDto struct {
	ID        uint   `json:"id" example:"1"`
	Email     string `json:"email" example:"john.doe@gmail.com"`
	Gsm       string `json:"gsm" example:"555-555-5555"`
	FirstName string `json:"firstName" example:"John"`
	LastName  string `json:"lastName" example:"Doe"`
}

// / Request for user list
type UserResponse struct {
	ApiResponse
	User UserDto `json:"user"`
}

// / Request for user list
type UserUpdateRequest struct {
	UserDto
}

// / Response for user update
type UserUpdateResponse struct {
	ApiResponse
	User UserDto `json:"user"`
}

// / Request for user deletion
type UserDeleteRequest struct {
	ID uint `json:"id" example:"1"`
}

// / Response for user deletion
type UserDeleteResponse struct {
	ApiResponse
}

// / Credential DTO
type CredentialDto struct {
	ApiResponse

	ID       uint   `json:"id" example:"1"`
	UserId   uint   `json:"userId" example:"1"`
	Email    string `json:"email" example:"john.doe@gmail.com"`
	Password string `json:"password" example:"123456"`
}

// / Request for user list
type UserListResponse struct {
	ApiResponse

	/// List of users
	Users []UserDto `json:"users"`
}

// / Request for user creation
type UserCreateRequest struct {

	/// User email
	Email string `json:"email" example:"john.doe@gmail.com"`
	/// User password
	Password string `json:"password" example:"123456"`
	/// User GSM
	Gsm string `json:"gsm" example:"555-555-5555"`
	/// User first name
	FirstName string `json:"firstName" example:"John"`
	/// User last name
	LastName string `json:"lastName" example:"Doe"`
}

// / Response for user creation
type UserCreateResponse struct {
	ApiResponse

	User       UserDto       `json:"user"`
	Credential CredentialDto `json:"credential"`
}

// Models for auth.go
type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// / Auth response
type AuthResponse struct {
	ApiResponse
}

// / Auth response data
type AuthResponseData struct {
	/// User ID
	UserId uint `json:"user_id"`
	/// Token
	Token string `json:"token"`
}

// / Request for token expiration
type ExpireTokenRequest struct {
	Token string `json:"token"`
}

// / Response for token expiration
type ExpireTokenResponse struct {
	ApiResponse
}

// / Request for token exists
type TokenExistsRequest struct {
	Token string `json:"token"`
}

// / Response for token exists
type TokenExistsResponse struct {
	ApiResponse

	/// Token exists or not
	Exists bool `json:"exists"`
}

// / Request for current tokens
type CurrentTokensResponse struct {
	ApiResponse
}

// / Request for all tokens
type ExpireAllTokensResponse struct {
	ApiResponse
}
