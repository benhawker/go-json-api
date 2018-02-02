package models

import (
	"github.com/revel/revel"
	"time"
)

type Property struct {
	UUID          string `sql:"type:uuid;primary_key" json:"uuid"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	PostalCode    string `json:"postal_code"`
	PricePerMonth int    `json:"price_per_month"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Property) Validate(v *revel.Validation) {
	v.Required(p.UUID)
	v.Required(p.Title)
	v.Required(p.Description)
	v.Required(p.PostalCode)
	v.Required(p.PricePerMonth)
}
