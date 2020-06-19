//
//  go-unit-test-sql
//
//  Copyright © 2020. All rights reserved.
//

package mysql

import (
	"database/sql"
	"log"
	"testing"

	r "github.com/moemoe89/go-unit-test-sql/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var u = &r.UserModel{
	ID:    uuid.New().String(),
	Name:  "Momo",
	Email: "momo@mail.com",
	Phone: "08123456789",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestFindByID(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "SELECT id, name, email, phone FROM users WHERE id = \\?"

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone"}).
		AddRow(u.ID, u.Name, u.Email, u.Phone)

	mock.ExpectQuery(query).WithArgs(u.ID).WillReturnRows(rows)

	user, err := repo.FindByID(u.ID)
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestFindByIDError(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "SELECT id, name, email, phone FROM user WHERE id = \\?"

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone"})

	mock.ExpectQuery(query).WithArgs(u.ID).WillReturnRows(rows)

	user, err := repo.FindByID(u.ID)
	assert.Empty(t, user)
	assert.Error(t, err)
}

func TestFind(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "SELECT id, name, email, phone FROM users"

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone"}).
		AddRow(u.ID, u.Name, u.Email, u.Phone)

	mock.ExpectQuery(query).WillReturnRows(rows)

	users, err := repo.Find()
	assert.NotEmpty(t, users)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
}

func TestCreate(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "INSERT INTO users \\(id, name, email, phone\\) VALUES \\(\\?, \\?, \\?, \\?\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.ID, u.Name, u.Email, u.Phone).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Create(u)
	assert.NoError(t, err)
}

func TestCreateError(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "INSERT INTO user \\(id, name, email, phone\\) VALUES \\(\\?, \\?, \\?, \\?\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.ID, u.Name, u.Email, u.Phone).WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.Create(u)
	assert.Error(t, err)
}

func TestUpdate(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "UPDATE users SET name = \\?, email = \\?, phone = \\? WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.Name, u.Email, u.Phone, u.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Update(u)
	assert.NoError(t, err)
}

func TestUpdateErr(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "UPDATE user SET name = \\?, email = \\?, phone = \\? WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.Name, u.Email, u.Phone, u.ID).WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.Update(u)
	assert.Error(t, err)
}

func TestDelete(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "DELETE FROM users WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Delete(u.ID)
	assert.NoError(t, err)
}

func TestDeleteError(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "DELETE FROM user WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.ID).WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.Delete(u.ID)
	assert.Error(t, err)
}
