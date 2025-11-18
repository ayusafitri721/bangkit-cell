package model

import (
	"database/sql"
	"time"
)

type TransactionDetail struct {
	ID                      int             `json:"id"`
	TransactionID           int             `json:"transaction_id"`
	DeviceServiceVariantID  int             `json:"device_service_variant_id"`
	Harga                   float64         `json:"harga"`
	HargaModal              sql.NullFloat64 `json:"harga_modal"`
	CreatedAt               time.Time       `json:"created_at"`
	UpdatedAt               time.Time       `json:"updated_at"`
}