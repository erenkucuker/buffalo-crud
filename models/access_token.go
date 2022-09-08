package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// AccessToken is used by pop to map your access_tokens database table to your go code.
type AccessToken struct {
	ID          int64     `json:"id" db:"id"`
	UserID      uuid.UUID `json:"user_id" db:"user_id"`
	AccessToken string    `json:"access_token" db:"access_token"`
	ExpiresAt   time.Time `json:"expires_at" db:"expires_at"`
	User        *User     `json:"user,omitempty" belongs_to:"user"`
}

// String is not required by pop and may be deleted
func (a AccessToken) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// AccessTokens is not required by pop and may be deleted
type AccessTokens []AccessToken

// String is not required by pop and may be deleted
func (a AccessTokens) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *AccessToken) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *AccessToken) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *AccessToken) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
