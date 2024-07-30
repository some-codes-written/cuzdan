package common

import (
	"time"

	"gorm.io/gorm"
)

// ModelBase is the base model for all entities
type ModelBase struct {
	gorm.Model

	Hash string `bson:"hash" json:"hash" gorm:"unique, not null;type:char(36)" example:"UUID"` // UUID
}

// AuditModel is the base model for all entities
type AuditModel struct {
	CreatedBy uint `json:"created_by" example:"1" gorm:"default:null;"` // Created by user id
	UpdatedBy uint `json:"updated_by" example:"1" gorm:"default:null;"` // Updated by user id
}

// SoftDeleteModel is the base model for all entities
type SoftDeleteModel struct {
	IsDeleted bool `json:"is_deleted" example:"false" gorm:"default:false;"` // Is deleted
	DeletedBy uint `json:"deleted_by" example:"1" gorm:"default:null;"`      // Deleted by user id
}

// EntityModel is the base model for all entities
type EntityModel struct {
	ModelBase
	AuditModel
	SoftDeleteModel
}

// IsTransient returns true if the model is transient (not persisted in the database)
func (s *EntityModel) IsTransient() bool {
	return s.ID == 0 || s.Hash == ""
}

// IsPersisted returns true if the model is persisted in the database
func (s *EntityModel) IsPersisted() bool {
	return s.ID != 0 && s.Hash != ""
}

func (s *EntityModel) SetHash() {
	// If the model is transient, generate a new hash
	if s.IsTransient() {
		s.Hash = GenerateToken()
	}
}

func (s *EntityModel) SetDefaultsviaCreation() {
	s.SetHash()
	s.CreatedAt = time.Now()
	s.IsDeleted = false
}

func (s *EntityModel) SetDefaultsviaDeletion() {
	s.IsDeleted = true
}

func (s *EntityModel) SetDefaultsviaUpdation() {
	s.UpdatedAt = time.Now()
}

func (s *EntityModel) ChangeEntityHash() {
	//TODO: save history for old hash
	s.Hash = GenerateToken()
}
