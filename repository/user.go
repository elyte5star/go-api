package repository

import (
	"time"

	"github.com/api/service/dbutils/schema"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserQueries struct {
	*sqlx.DB
}

type GetUserResponse struct {
	Userid           uuid.UUID           `json:"userid"`
	LastModifiedAt   time.Time           `json:"lastModifiedAt"`
	CreatedAt        time.Time           `json:"createdAt"`
	Username         string              `json:"username"`
	Email            string              `json:"email"`
	AccountNonLocked bool                `json:"account_not_locked"`
	Admin            bool                `json:"admin"`
	Enabled          bool                `json:"enabled"`
	IsUsing2FA       bool                `json:"isUsing2FA"`
	Telephone        string              `json:"telephone"`
	LockTime         time.Time           `json:"lockTime"`
	Address          *schema.UserAddress `json:"address"`
	Bookings         []*schema.Booking   `json:"bookings"`
}

type GetUsersResponse struct {
	Users []GetUserResponse `json:"users"`
}

func (q *UserQueries) GetUserById(userid uuid.UUID) (*GetUserResponse, error) {

	var user schema.User
	// Define query string.
	query := `SELECT * FROM users WHERE userid=?`

	// Send query to database.
	err := q.Get(&user, query, userid)
	if err != nil {
		// Return empty object and error.
		return &GetUserResponse{}, err
	}
	result := &GetUserResponse{Userid: user.Userid,
		LastModifiedAt:   user.AuditInfo.LastModifiedAt,
		CreatedAt:        user.AuditInfo.CreatedAt,
		Username:         user.UserName,
		Email:            user.Email,
		AccountNonLocked: user.AccountNonLocked,
		Admin:            user.Admin,
		IsUsing2FA:       user.IsUsing2FA,
		Enabled:          user.Enabled,
		Telephone:        user.Telephone,
		LockTime:         user.LockTime,
	}
	return result, nil
}

func (q *UserQueries) GetUsers() (GetUsersResponse, error) {
	// Define users variable.

	var users GetUsersResponse
	// Define query string.
	query := `SELECT * FROM users`

	// Send query to database.
	err := q.Get(&users, query)
	if err != nil {
		// Return empty object and error.
		return users, err
	}

	// Return query result.
	return users, nil
}

// Createuser method for creating User by given User object.
func (q *UserQueries) CreateUser(user *schema.User) error {
	// Define query string.
	query := `INSERT INTO users (userid,username,password,email,telephone,lockTime,auditInfo)
	 VALUES (:userid,:username,:password,:email,:telephone,:lockTime,:auditInfo)`

	// Send query to database.
	_, err := q.NamedExec(query, user)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// UpdateUser method for updating user by given User object.
func (q *UserQueries) UpdateUser(userid uuid.UUID, user *schema.User) error {
	// Define query string.
	query := `UPDATE users SET lastModifiedAt = $2, LastModifiedBy = $3, telephone = $4, email = $5, address = $6 WHERE userid=?`

	// Send query to database.
	_, err := q.Exec(query, userid, user.AuditInfo.LastModifiedAt, user.AuditInfo.LastModifiedBy, user.Telephone, user.Email)
	if err != nil {
		// Return only error.
		return err
	}
	// This query returns nothing.
	return nil
}

// DeleteUser method for delete user by given ID.
func (q *UserQueries) DeleteUser(userid uuid.UUID) error {
	// Define query string.
	query := `DELETE FROM books WHERE id = $1`

	// Send query to database.
	_, err := q.Exec(query, userid)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
