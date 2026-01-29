package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"laundry/internal/model"
	"laundry/internal/repository"
	"laundry/internal/service"
)

type TransactionHandler struct {
	txService *service.TransactionService
}

func NewTransactionHandler(
	txService *service.TransactionService,
) *TransactionHandler {
	return &TransactionHandler{
		txService: txService,
	}
}

func (h *TransactionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email     string   `json:"email"`
		Nama      string   `json:"nama"`
		Alamat    *string  `json:"alamat"`
		Tanggal   string   `json:"tanggal"` // yyyy-mm-dd
		QtyIn     int      `json:"qty_in"`
		QtyOut    int      `json:"qty_out"`
		Deskripsi *string  `json:"deskripsi"`
		AmountIn  float64  `json:"amount_in"`
		AmountOut float64  `json:"amount_out"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	tanggal, err := time.Parse("2006-01-02", req.Tanggal)
	if err != nil {
		http.Error(w, "invalid tanggal format (yyyy-mm-dd)", http.StatusBadRequest)
		return
	}

	tx := &model.Transaction{
		Email:      req.Email,
		Nama:       req.Nama,
		Alamat:     req.Alamat,
		Tanggal:    tanggal,
		QtyIn:      req.QtyIn,
		QtyOut:     req.QtyOut,
		Deskripsi:  req.Deskripsi,
		AmountIn:   req.AmountIn,
		AmountOut:  req.AmountOut,
	}

	if err := h.txService.Create(r.Context(), tx); err != nil {
		h.handleServiceError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"id": tx.ID,
	})
}

func (h *TransactionHandler) GetByEmailAndDate(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	tanggalStr := r.URL.Query().Get("tanggal")

	if email == "" || tanggalStr == "" {
		http.Error(w, "email and tanggal are required", http.StatusBadRequest)
		return
	}

	_, err := time.Parse("2006-01-02", tanggalStr)
	if err != nil {
		http.Error(w, "invalid tanggal format (yyyy-mm-dd)", http.StatusBadRequest)
		return
	}

	log.Printf("[DEBUG] Fetching Transactions for Email: %s, Date: %s", email, tanggalStr)

	tx, err := h.txService.ListByEmailAndDate(r.Context(), email, tanggalStr)
	if err != nil {
		if errors.Is(err, repository.ErrTransactionNotFound) {
			// Kirim data kosong daripada 404 agar UI tidak error
			emptyTx := map[string]any{
				"email":      email,
				"tanggal":    tanggalStr,
				"qty_in":     0,
				"qty_out":    0,
				"amount_in":  0,
				"amount_out": 0,
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(emptyTx)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(tx)
}


func (h *TransactionHandler) handleServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, service.ErrInvalidQty),
		errors.Is(err, service.ErrNegativeQty),
		errors.Is(err, service.ErrInvalidAmount),
		errors.Is(err, service.ErrFutureDate):

		http.Error(w, err.Error(), http.StatusBadRequest)

	default:
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
