//
//  go-unit-test-sql
//
//  Copyright © 2020. All rights reserved.
//

package postgres

import (
	"context"
	"database/sql"
	"time"

	repo "github.com/moemoe89/go-unit-test-sql/repository"

	_ "github.com/lib/pq"
)

// repository represent the repository model
type repository struct {
	db *sql.DB
}

// NewRepository will create a variable that represent the Repository struct
func NewRepository(dialect, dsn string, idleConn, maxConn int) (repo.Repository, error) {
	db, err := sql.Open(dialect, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(idleConn)
	db.SetMaxOpenConns(maxConn)

	return &repository{db}, nil
}

// Close attaches the provider and close the connection
func (r *repository) Close() {
	r.db.Close()
}

// FindByID attaches the user repository and find data based on id
func (r *repository) FindByID(id string) (*repo.UserModel, error) {
	user := new(repo.UserModel)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, "SELECT id, name, email, phone FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Phone)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Find attaches the user repository and find all data
func (r *repository) Find() ([]*repo.UserModel, error) {
	users := make([]*repo.UserModel, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, "SELECT id, name, email, phone FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := new(repo.UserModel)
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Phone,
		)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Create attaches the user repository and creating the data
func (r *repository) Create(user *repo.UserModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "INSERT INTO users (id, name, email, phone) VALUES ($1, $2, $3, $4)"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.ID, user.Name, user.Email, user.Phone)
	return err
}

// Update attaches the user repository and update data based on id
func (r *repository) Update(user *repo.UserModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "UPDATE users SET name = $1, email = $2, phone = $3 WHERE id = $4"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.Name, user.Email, user.Phone, user.ID)
	return err
}

// Delete attaches the user repository and delete data based on id
func (r *repository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "DELETE FROM users WHERE id = $1"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	return err
}
