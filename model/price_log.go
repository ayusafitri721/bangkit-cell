package model

import (
	"database/sql"
	"time"
)

type PriceLog struct {
	ID                      int             `json:"id"`
	DeviceServiceVariantID  int             `json:"device_service_variant_id"`
	UserID                  sql.NullInt64   `json:"user_id"`
	OldHargaMin             sql.NullFloat64 `json:"old_harga_min"`
	OldHargaMax             sql.NullFloat64 `json:"old_harga_max"`
	NewHargaMin             sql.NullFloat64 `json:"new_harga_min"`
	NewHargaMax             sql.NullFloat64 `json:"new_harga_max"`
	TipePart                sql.NullString  `json:"tipe_part"`
	ChangedAt               time.Time       `json:"changed_at"`
}