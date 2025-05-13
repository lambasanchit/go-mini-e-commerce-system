package service

import (
	"fmt"
	"user-service/model"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func RegisterUser(username, password string) (*model.User, error) {
	// Prevent duplicate registration
	if _, exists := model.UserStore[username]; exists {
		return nil, fmt.Errorf("user already exists")
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("error registering user: %w", err)
	}

	user := &model.User{
		ID:       fmt.Sprintf("%d", len(model.UserStore)+1),
		Username: username,
		Password: hashedPassword,
	}

	model.UserStore[username] = user
	return user, nil
}

func LoginUser(username, password string) bool {
	user, exists := model.UserStore[username]
	if !exists {
		return false
	}
	return CheckPasswordHash(password, user.Password)
}
