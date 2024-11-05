package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	track "github.com/Namdar1Ibrakhim/student-track-system"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/constants"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/dto"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user track.User, role constants.Role) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user, role)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil

}

func (s *AuthService) GetUser(userId int) (dto.UserResponse, error) {
	user, error := s.repo.FindByID(userId)
	if error != nil {
		return dto.UserResponse{}, error
	}

	return user, nil
}

func (s *AuthService) GetAllUsers() ([]dto.GetAllUsersResponse, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *AuthService) UpdateUser(userId int, input dto.UpdateUser) error {
	return s.repo.UpdateUser(userId, input)
}

func (s *AuthService) DeleteUser(userId int) error {
	return s.repo.DeleteUser(userId)
}

func (s *AuthService) EditPassword(userId int, oldPassword, newPassword string, isAdmin bool) error {
	if !isAdmin {
		currentUser, err := s.repo.GetPasswordHashById(userId)
		if err != nil {
			return fmt.Errorf("user not found")
		}

		if currentUser.Password_hash != generatePasswordHash(oldPassword) {
			return fmt.Errorf("incorrect old password")
		}
	}

	newPasswordHash := generatePasswordHash(newPassword)
	return s.repo.EditPassword(userId, newPasswordHash)
}
