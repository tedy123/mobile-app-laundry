package service

import (
	"context"
	"errors"
	"time"

	"laundry/internal/model"
	"laundry/internal/repository"
)

var (
	ErrInvalidQty       = errors.New("either qty_in or qty_out must be filled")
	ErrNegativeQty      = errors.New("qty cannot be negative")
	ErrInvalidAmount    = errors.New("either amount_in or amount_out must be greater than zero")
	ErrFutureDate       = errors.New("transaction date cannot be in the future")
)

type TransactionService struct {
	txRepo *repository.TransactionRepository
}

func NewTransactionService(
	txRepo *repository.TransactionRepository,
) *TransactionService {
	return &TransactionService{
		txRepo: txRepo,
	}
}

func (s *TransactionService) Create(
	ctx context.Context,
	tx *model.Transaction,
) error {

	if tx.QtyIn < 0 || tx.QtyOut < 0 {
		return ErrNegativeQty;
	}

	if (tx.QtyIn > 0 && tx.QtyOut > 0) ||
		(tx.QtyIn == 0 && tx.QtyOut == 0) {
		return ErrInvalidQty;
	}

	if (tx.AmountIn > 0 && tx.AmountOut > 0) ||
		(tx.AmountIn == 0 && tx.AmountOut == 0) {
		return ErrInvalidAmount;
	}

	if tx.Tanggal.After(time.Now()) {
		return ErrFutureDate;
	}

	return s.txRepo.Create(ctx, tx)
}

func (s *TransactionService) ListByEmailAndDate(
	ctx context.Context,
	email string,
	tanggal string,
) (*model.Transaction, error) {

	return s.txRepo.FindByEmailAndDate(ctx, email, tanggal)
}
