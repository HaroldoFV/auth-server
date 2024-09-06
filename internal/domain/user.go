package domain

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strings"
	"time"
)

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type User struct {
	id        string
	name      string
	email     string
	password  string
	status    string
	createdAt time.Time
}

func NewUser(name, email, password string) (*User, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name cannot be empty or consist only of whitespace characters")
	}

	if len(password) < 8 {
		return nil, errors.New("password must be at least 8 characters long")
	}

	if strings.TrimSpace(password) == "" {
		return nil, errors.New("password cannot consist only of whitespace characters")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(password)), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	user := &User{
		id:        uuid.New().String(),
		name:      strings.TrimSpace(name),
		email:     strings.TrimSpace(email),
		password:  string(hash),
		status:    ENABLED,
		createdAt: time.Now(),
	}
	if err := user.IsValid(); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) IsValid() error {
	if u.id == "" {
		return errors.New("invalid id")
	}
	if u.name == "" {
		return errors.New("name cannot be empty")
	}
	if len(u.name) > 100 {
		return errors.New("name cannot be longer than 100 characters")
	}

	if u.email == "" {
		return errors.New("email cannot be empty")
	}
	if !isValidEmail(u.email) {
		return errors.New("invalid email format")
	}
	return nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
	return err == nil
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func (u *User) Disable() error {
	u.status = DISABLED
	err := u.IsValid()
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Enabled() error {
	u.status = ENABLED
	err := u.IsValid()
	if err != nil {
		return err
	}
	return nil
}

func (u *User) ChangePassword(oldPassword, newPassword string) error {
	if !u.ValidatePassword(oldPassword) {
		return errors.New("invalid old password")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.password = string(hashedPassword)
	return nil
}

func (u *User) GetID() string {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetStatus() string {
	return u.status
}

func (u *User) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *User) SetID(id string) {
	u.id = id
}

func (u *User) SetCreatedAt(date time.Time) {
	u.createdAt = date
}
