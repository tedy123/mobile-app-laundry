package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"laundry/internal/model"
	"laundry/internal/repository"
	"laundry/internal/utils"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

type GoogleTokenInfo struct {
	Email string `json:"email"`
	Sub   string `json:"sub"` // Google UID
	Aud   string `json:"aud"` // Client ID
	Exp   string `json:"exp"`
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not found")
	}

	if user.Password == nil {
		return "", errors.New("password login not allowed")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(*user.Password),
		[]byte(password),
	); err != nil {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateToken(user.Email, string(user.Role))
}

func (s *AuthService) GoogleLogin(ctx context.Context, idToken string) (string, string, error) {
	// Verifikasi token manual menggunakan API Google
	// Ini tidak memerlukan file Firebase JSON
	url := fmt.Sprintf("https://oauth2.googleapis.com/tokeninfo?id_token=%s", idToken)
	
	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", errors.New("invalid google token")
	}

	var info GoogleTokenInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return "", "", err
	}

	email := info.Email
	if email == "" {
		return "", "", errors.New("email not found in token")
	}

	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}

	if user == nil {
		// Auto-register
		user = &model.User{
			ID:        uuid.New().String(),
			Email:     email,
			GoogleUID: &info.Sub,
			Role:      model.Cashier,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := s.userRepo.Create(ctx, user); err != nil {
			return "", "", err
		}
	} else if user.GoogleUID == nil {
		if err := s.userRepo.UpdateGoogleUID(ctx, email, info.Sub); err != nil {
			return "", "", err
		}
	}

	accessToken, err := utils.GenerateToken(user.Email, string(user.Role))
	return accessToken, user.Email, err
}
