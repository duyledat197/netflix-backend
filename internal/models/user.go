package models

import (
	"github.com/duyledat197/netfix-backend/common/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

//User ...
type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Email        string             `bson:"email"`
	NiceName     string             `bson:"niceName"`
	Password     string             `bson:"password"`
	FullName     string             `bson:"fullName"`
	Role         enum.UserRole      `bson:"role"`
	PostCount    int64              `bson:"postCount"`
	CommentCount int64              `bson:"commentCount"`
}

// HashPassword : hash password using crypto
func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

// IsCorrectPassword : check password with passwordhash
func (u *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
