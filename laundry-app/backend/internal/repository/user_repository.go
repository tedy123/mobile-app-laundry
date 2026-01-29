package repository

import (
	"context"
	"database/sql"
	"laundry/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, email, password, google_uid, role, created_at, updated_at
		FROM users WHERE email=$1
	`, email)

	var u model.User
	err := row.Scan(
		&u.ID,
		&u.Email,
		&u.Password,
		&u.GoogleUID,
		&u.Role,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &u, err
}

func (r *UserRepository) Create(ctx context.Context, u *model.User) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO users (id, email, password, google_uid, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, u.ID, u.Email, u.Password, u.GoogleUID, u.Role, u.CreatedAt, u.UpdatedAt)
	return err
}

func (r *UserRepository) UpdateGoogleUID(ctx context.Context, email, googleUID string) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE users SET google_uid = $1, updated_at = NOW() WHERE email = $2
	`, googleUID, email)
	return err
}
