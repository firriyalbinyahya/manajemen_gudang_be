package service

import (
	"errors"
	"log"
	"manajemen_gudang_be/entity"
	"manajemen_gudang_be/repository"
	"manajemen_gudang_be/utils/auth"
	"manajemen_gudang_be/utils/response"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (as *AuthService) Register(req *entity.RegisterRequest) error {
	_, err := as.UserRepository.GetUserByUsername(req.Username)
	if err == nil {
		return errors.New("username already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := &entity.User{
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
	}

	err = as.UserRepository.CreateUser(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (as *AuthService) Login(req *entity.LoginRequest) (*entity.LoginResponse, error) {
	user, err := as.UserRepository.GetUserByUsername(req.Username)
	if err != nil {
		return &entity.LoginResponse{}, response.NewCustomError(http.StatusUnauthorized, "username or password information does not match", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		log.Printf("Error compare password: %v", err)
		return &entity.LoginResponse{}, response.NewCustomError(http.StatusUnauthorized, "username or password information does not match", err)
	}

	// generate token
	accessToken, accessExp, err := auth.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &entity.LoginResponse{
		Username:      user.Username,
		AccessToken:   accessToken,
		AccessExpired: accessExp,
	}, nil
}
