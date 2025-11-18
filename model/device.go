package model

import "time"

type Device struct {
	ID        int       `json:"id"`
	BrandID   int       `json:"brand_id"`
	Model     string    `json:"model"`
	Tipe      string    `json:"tipe"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}