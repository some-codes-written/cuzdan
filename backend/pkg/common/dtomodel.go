package common

type DTOModel struct {
	DTOModelWithIdentifier
	DTOModelWithTimestamp
}

type DTOModelWithIdentifier struct {
	ID string `json:"id"`
}

type DTOModelWithTimestamp struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
