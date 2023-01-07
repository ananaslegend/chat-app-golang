package repository

import (
	"chat-app-golang/internal/domain"
	"context"
)

type User interface {
	Get(ctx context.Context, user domain.User) (*domain.User, error)
	Insert(ctx context.Context, user domain.User) error
	Update(ctx context.Context, user domain.User)
	Archive(ctx context.Context, user domain.User)
	Delete(ctx context.Context, user domain.User)
	GetUsersByCredentials(ctx context.Context, user domain.User) (*[]domain.User, error)
}
