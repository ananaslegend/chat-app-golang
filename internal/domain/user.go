package domain

type User struct {
	ID           *string    `bson:"_id,omitempty"`
	Username     *string    `bson:"username,omitempty"`
	Email        *string    `bson:"email,omitempty"`
	PasswordHash *string    `bson:"passwordHash,omitempty"`
	State        *UserState `bson:"state,omitempty"`
}

type UserState int

const (
	Archived UserState = iota
	Active
)
