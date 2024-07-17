package repository

import (
	"github.com/api/repository/schema"
	"github.com/jmoiron/sqlx"
)

type AuthQueries struct {
	*sqlx.DB
}

func (auth *AuthQueries) FindByUsername(username string) (schema.User, error) {
	user := schema.User{}
	query := `SELECT * FROM users WHERE username=?`
	// Send query to database.
	err := auth.Get(&user, query, username)
	if err != nil {
		// Return empty object and error.
		return user, err
	}
	return user, nil
}

func (auth *AuthQueries) FindByEmail(email string) (schema.User, error) {
	user := schema.User{}
	query := `SELECT * FROM users WHERE email=?`
	// Send query to database.
	err := auth.Get(&user, query, email)
	if err != nil {
		// Return empty object and error.
		return user, err
	}
	return user, nil
}
