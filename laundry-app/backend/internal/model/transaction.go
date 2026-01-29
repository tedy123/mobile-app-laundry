package model

import (
	"time"
)

type Transaction struct {
	ID         string    `json:"id"`
	Email      string    `json:"email"`
	Nama       string    `json:"nama"`
	Alamat     *string   `json:"alamat"`
	Tanggal    time.Time `json:"tanggal"`
	QtyIn      int       `json:"qty_in"`
	QtyOut     int       `json:"qty_out"`
	Deskripsi  *string   `json:"deskripsi"`
	AmountIn   float64   `json:"amount_in"`
	AmountOut  float64   `json:"amount_out"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
