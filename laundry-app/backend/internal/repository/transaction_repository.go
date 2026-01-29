package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"laundry/internal/model"
)

var ErrTransactionNotFound = errors.New("transaction not found")

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (r *TransactionRepository) Create(
	ctx context.Context,
	tx *model.Transaction,
) error {
	query := `
		INSERT INTO transactions (
			email, nama, alamat, tanggal, qty_in, qty_out, deskripsi, amount_in, amount_out
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		RETURNING id
	`
	return r.db.QueryRowContext(
		ctx, query,
		tx.Email, tx.Nama, tx.Alamat, tx.Tanggal,
		tx.QtyIn, tx.QtyOut, tx.Deskripsi, tx.AmountIn, tx.AmountOut,
	).Scan(&tx.ID)
}

func (r *TransactionRepository) FindByEmailAndDate(
	ctx context.Context,
	email string,
	tanggal string,
) (*model.Transaction, error) {

	row := r.db.QueryRowContext(ctx, `
		SELECT
			COALESCE(SUM(qty_in), 0),
			COALESCE(SUM(qty_out), 0),
			COALESCE(SUM(amount_in), 0),
			COALESCE(SUM(amount_out), 0)
		FROM transactions
		WHERE TRIM(LOWER(email)) = TRIM(LOWER($1)) 
		  AND (tanggal::text LIKE $2 || '%')
	`, email, tanggal)

	var s model.Transaction
	s.Email = email

	if err := row.Scan(&s.QtyIn, &s.QtyOut, &s.AmountIn, &s.AmountOut); err != nil {
		return nil, err
	}

	return &s, nil
}
