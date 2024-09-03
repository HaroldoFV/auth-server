package database

import (
	"auth-server/internal/domain"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	fmt.Printf("Creating UserRepository with db: %v\n", db)
	return &UserRepository{Db: db}
}

func (r *UserRepository) Create(user *domain.User) error {
	fmt.Printf("user: %v", user)
	stmt, err := r.Db.Prepare("INSERT INTO users (id, name, email, password, status, created_at) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.GetID(), user.GetName(), user.GetEmail(), user.GetPassword(), user.GetStatus(), user.GetCreatedAt())
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	query := "SELECT id, name, email, password, status, created_at FROM users WHERE email = $1"

	row := r.Db.QueryRow(query, email)

	var user *domain.User
	var idStr, name, emailStr, password, status string
	var createdAt time.Time

	err := row.Scan(&idStr, &name, &emailStr, &password, &status, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, err
	}

	user, err = domain.NewUser(name, email, password)
	if err != nil {
		return nil, err
	}

	user.SetID(idStr)
	user.SetCreatedAt(createdAt)

	if status == domain.ENABLED {
		err = user.Enabled()
	} else {
		err = user.Disable()
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetByName(name string) (*domain.User, error) {
	query := "SELECT id, name, email, password, status, created_at FROM users WHERE name = $1"

	row := r.Db.QueryRow(query, name)

	var user *domain.User
	var idStr, nameStr, email, password, status string
	var createdAt time.Time

	err := row.Scan(&idStr, &nameStr, &email, &password, &status, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with name %s not found", name)
		}
		return nil, err
	}

	user, err = domain.NewUser(name, email, password)
	if err != nil {
		return nil, err
	}

	user.SetID(idStr)
	user.SetCreatedAt(createdAt)

	if status == domain.ENABLED {
		err = user.Enabled()
	} else {
		err = user.Disable()
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}
