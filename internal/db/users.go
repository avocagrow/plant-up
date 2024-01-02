package db

import (
	"context"
	"errors"
	"time"
)

type User struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	DisplayName string    `json:"display_name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	IdpID       string    `json:"-"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

func (i *DbInstance) CreateUser(ctx context.Context, u *User) (*User, error) {
	q := `INSERT INTO "users" ("name","display_name","email","phone","idp_id")
    VALUES ($1,$2,$3,$4,$5) returning id;
    `
	var lastInsertID string
	row := i.Pool.QueryRow(ctx, q, u.Name, u.DisplayName, u.Email, u.Phone)
	if err := row.Scan(&lastInsertID); err != nil {
		return nil, err
	}
	u.ID = lastInsertID
	return u, nil
}

func (i *DbInstance) UpdateUser(ctx context.Context, u *User) (*User, error) {
	return nil, errors.New("not implemented")
}

func (i *DbInstance) DeleteUser(ctx context.Context, id string) error {
	return errors.New("not implemented")
}

func (i *DbInstance) GetUserByID(ctx context.Context, id string) (*User, error) {
	return nil, errors.New("not implemented")
}

func (i *DbInstance) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	return nil, errors.New("not implemented")
}

func (i *DbInstance) GetUserByPhone(ctx context.Context, phone string) (*User, error) {
	return nil, errors.New("not implemented")
}
