package user

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int
	Username  string `gorm:"unique"`
	Iat       int
	Exp       int
	Salt      string
	Hash      string
	Role      string    `gorm:"type:enum('admin', 'user');default:'user'"`
	CreatedAt time.Time `gorm:"<-:create"`
}

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrPasswordIncorrect = errors.New("password incorrect")
)

func GetUser(r Repository, username, password string) (*User, error) {
	// find user by username if exists in db check password hash
	user := &User{}
	var err error
	user, err = r.FindByUsername(username)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func Create(r Repository, u *User) error {
	return r.Create(u)
}

func Update(r Repository, u *User) error {
	return r.Update(u)
}

func Delete(r Repository, id int) error {
	return r.Delete(id)
}

func FindByID(r Repository, id int) (*User, error) {
	return r.FindByID(id)
}

func FindAll(r Repository) (*[]User, error) {
	return r.FindAll()
}
