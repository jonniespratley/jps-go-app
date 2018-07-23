package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Theme struct {
	ID              uuid.UUID    `json:"id" db:"id"`
	CreatedAt       time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at" db:"updated_at"`
	URI             string       `json:"uri" db:"uri"`
	BaseUri         string       `json:"base_uri" db:"base_uri"`
	Main            string       `json:"main" db:"main"`
	Error           string       `json:"error" db:"error"`
	ErrorChromeless string       `json:"error_chromeless" db:"error_chromeless"`
	Options         nulls.String `json:"options" db:"options"`
	DisplayName     string       `json:"display_name" db:"display_name"`
	DemoLink        nulls.String `json:"demo_link" db:"demo_link"`
	Description     nulls.String `json:"description" db:"description"`
}

// String is not required by pop and may be deleted
func (t Theme) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Themes is not required by pop and may be deleted
type Themes []Theme

// String is not required by pop and may be deleted
func (t Themes) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Theme) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: t.URI, Name: "URI"},
		&validators.StringIsPresent{Field: t.BaseUri, Name: "BaseUri"},
		&validators.StringIsPresent{Field: t.Main, Name: "Main"},
		&validators.StringIsPresent{Field: t.Error, Name: "Error"},
		&validators.StringIsPresent{Field: t.ErrorChromeless, Name: "ErrorChromeless"},
		&validators.StringIsPresent{Field: t.DisplayName, Name: "DisplayName"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Theme) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Theme) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
