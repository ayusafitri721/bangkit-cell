package model

import "time"

type Brand struct {
	ID          int       `json:"id"`
	Nama        string    `json:"nama"`
	NegaraAsal  string    `json:"negara_asal"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}