package domain

type User struct {
	ID           string `bson:"_id,omitempty"`
	Username     string `bson:"username"`
	Email        string `bson:"email"`
	PasswordHash string `bson:"passwordHash"`
	State        int    `bson:"state"`
}
