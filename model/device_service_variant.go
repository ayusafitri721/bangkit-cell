package model

import (
	"database/sql"
	"time"
)

type DeviceServiceVariant struct {
	ID        int            `json:"id"`
	DeviceID  int            `json:"device_id"`
	ServiceID int            `json:"service_id"`
	TipePart  sql.NullString `json:"tipe_part"`
	HargaMin  float64        `json:"harga_min"`
	HargaMax  float64        `json:"harga_max"`
	Catatan   sql.NullString `json:"catatan"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}