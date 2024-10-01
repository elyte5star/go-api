package repository

import (
	"database/sql"
	"fmt"

	"github.com/api/repository/response"
	"github.com/api/repository/schema"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserQueries struct {
	*sqlx.DB
}

func (q *UserQueries) GetUserById(userid uuid.UUID) (schema.User, *response.GetUserAdressResponse, error) {
	user := schema.User{}
	address := &response.GetUserAdressResponse{}
	query := `SELECT u.*,a.fullName,a.streetAddress,a.country,a.state, a.zip
	 FROM users AS u LEFT JOIN address AS a on u.userid=a.userid WHERE u.userid=?`
	if err := q.QueryRow(query, userid).Scan(&user.Userid, &user.Username, &user.Password, &user.Email, &user.AccountNonLocked,
		&user.Admin, &user.Enabled, &user.IsUsing2FA, &user.Telephone, &user.Discount, &user.FailedAttempt,
		&user.LockTime, &user.AuditInfo, &address.FullName,
		&address.StreetAddress, &address.Country, &address.State, &address.Zip); err != nil {
		if err == sql.ErrNoRows {
			return user, address, fmt.Errorf("unknown userid : %d", userid)
		}
		return user, address, err
	}
	return user, address, nil
}

func (q *UserQueries) GetUserByUsername(username string) (schema.User, error) {

	user := schema.User{}
	// Define query string.
	query := `SELECT * FROM users WHERE username=?`

	// Send query to database.
	err := q.Get(&user, query, username)
	if err != nil {
		// Return empty object and error.
		return user, err
	}
	return user, nil
}

func (q *UserQueries) GetUserAddressById(userid uuid.UUID) (schema.UserAddress, error) {

	userAddress := schema.UserAddress{}
	// Define query string.
	query := `SELECT * FROM address WHERE userid=?`

	// Send query to database.
	err := q.Get(&userAddress, query, userid)
	if err != nil {
		// Return empty object and error.
		return userAddress, err
	}
	return userAddress, nil
}

// CreateUserAdress method for creating UserAddress by given UserAddress object.
func (q *UserQueries) CreateUserAdress(address *schema.UserAddress) error {
	// Define query string.
	query := `INSERT INTO address (userid,fullName,streetAddress,country,state,zip)
	 VALUES (:userid,:fullName,:streetAddress,:country,:state,:zip)`
	// Send query to database.
	_, err := q.NamedExec(query, address)
	if err != nil {
		// Return only error.
		return err
	}
	// This query returns nothing.
	return nil
}

// UpdateUserAdress method for creating UserAddress by given UserAddress object.
func (q *UserQueries) UpdateUserAdress(userid uuid.UUID, address *schema.UserAddress) error {
	// Define query string.
	query := `UPDATE address SET fullName=?,streetAddress=?,country=?,state=?,zip=? WHERE userid=?`
	_, err := q.Exec(query, address.FullName, address.StreetAddress, address.Country, address.State, address.Zip, userid)
	if err != nil {
		// Return only error.
		return err
	}
	// This query returns nothing.
	return nil
}

func (q *UserQueries) GetUsers() ([]schema.User, error) {
	// Define users variable.
	users := []schema.User{}
	// Define query string.
	query := `SELECT * FROM users`
	// Send query to database.
	err := q.Select(&users, query)
	if err != nil {
		// Return empty object and error.
		return users, err
	}
	return users, nil

}

// Createuser method for creating User by given User object.
func (q *UserQueries) CreateUser(user *schema.User) error {
	// Define query string.
	query := `INSERT INTO users (userid,username,password,email,telephone,lockTime,discount,
	accountNonLocked,admin,enabled,isUsing2FA,failedAttempt,auditInfo)
	 VALUES (:userid,:username,:password,:email,:telephone,:lockTime,:discount,:accountNonLocked,:admin,:enabled,:isUsing2FA,:failedAttempt,CONVERT(:auditInfo using utf8mb4))`
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
	query := `UPDATE users SET username=?,telephone=?, auditInfo=? WHERE userid=?`
	// Send query to database.
	_, err := q.Exec(query, user.Username, user.Telephone, user.AuditInfo, userid)
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
	query := `DELETE FROM users WHERE userid=?`
	// Send query to database.
	_, err := q.Exec(query, userid)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
