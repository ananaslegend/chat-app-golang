package repository

import "chat-app-golang/internal/domain"

type User interface {
	Get(user domain.User)
	Insert(user domain.User)
	Update(user domain.User)
	Archive(user domain.User)
	Delete(user domain.User)
}
