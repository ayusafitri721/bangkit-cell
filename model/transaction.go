package model

import (
	"database/sql"
	"time"
)

type Transaction struct {
	ID               int            `json:"id"`
	IDOperator       int            `json:"id_operator"`
	Status           string         `json:"status"`
	MetodePembayaran sql.NullString `json:"metode_pembayaran"`
	JumlahBayar      sql.NullFloat64 `json:"jumlah_bayar"`
	Kembalian        sql.NullFloat64 `json:"kembalian"`
	QrisReference    sql.NullString `json:"qris_reference"`
	Total            float64        `json:"total"`
	CustomerName     sql.NullString `json:"customer_name"`
	CustomerPhone    sql.NullString `json:"customer_phone"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
}