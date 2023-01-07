package service

import (
	"chat-app-golang/internal/db"
	"chat-app-golang/internal/db/repository"
	"chat-app-golang/internal/domain"
	"chat-app-golang/internal/dto"
	service "chat-app-golang/internal/service/errors"
	"chat-app-golang/pkg/pointer"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	userRepo repository.User
}

func NewAuth(userRepo repository.User) *Auth {
	return &Auth{userRepo: userRepo}
}

func (a Auth) Registration(regData dto.RegistrationUser) error {
	usersWithSameCredentials, err := a.userRepo.GetUsersByCredentials(context.TODO(), domain.User{Email: &regData.Email, Username: &regData.Username})
	if err == nil {
		for _, u := range *usersWithSameCredentials {
			if regData.Email == *u.Email {
				return service.EmailAlreadyExistErr
			}
			if regData.Username == *u.Username {
				return service.UsernameAlreadyExistErr
			}
		}
	} else if err != nil && !errors.Is(err, db.NoResultErr) {
		return err
	}

	hashedPass, err := hashPassword(regData.Password)
	if err != nil {
		return err
	}

	if err = a.userRepo.Insert(context.TODO(), domain.User{
		Email:        &regData.Email,
		Username:     &regData.Username,
		PasswordHash: &hashedPass,
		State:        (*domain.UserState)(pointer.ToInt(1)),
	}); err != nil {
		return err
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
