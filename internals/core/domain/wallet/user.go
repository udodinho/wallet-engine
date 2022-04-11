package wallet

import "time"

// User struct
type User struct {
	Reference       string    `json:"reference, omitempty" bson:"reference"`
	FirstName       string    `json:"first_name, omitempty" bson:"first_name"`
	LastName        string    `json:"last_name, omitempty" bson:"last_name"`
	Email           string    `json:"email, omitempty" bson:"email"`
	Currency        string    `json:"currency, omitempty" bson:"currency"`
	Password        string    `json:"-, omitempty" bson:"password"`
	HashedSecretKey string    `json:"-, omitempty" bson:"hashed_secret_key"`
	BVN             string    `json:"bvn, omitempty" bson:"bvn"`
	DateOfBirth     string    `json:"date_of_birth, omitempty" bson:"date_of_birth"`
	CreatedAt       time.Time `json:"created_at, omitempty" bson:"created_at"`
	IsActive        bool      `json:"is_active, omitempty" bson:"is_active"`
}
