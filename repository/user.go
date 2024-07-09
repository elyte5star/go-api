package repository

import (
	"github.com/api/repository/response"
	"github.com/api/service/dbutils/schema"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserQueries struct {
	*sqlx.DB
}

func (q *UserQueries) GetUserById(userid uuid.UUID) (schema.User, error) {

	user := schema.User{}
	// Define query string.
	query := `SELECT * FROM users WHERE userid=?`

	// Send query to database.
	err := q.Get(&user, query, userid)
	if err != nil {
		// Return empty object and error.
		return user, err
	}
	return user, nil
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

func (q *UserQueries) GetUserAddressById(userid uuid.UUID) (*response.GetUserAdressResponse, error) {

	userAddress := schema.UserAddress{}
	// Define query string.
	query := `SELECT * FROM address WHERE addressId=?`

	// Send query to database.
	err := q.Get(&userAddress, query, userid)
	if err != nil {
		// Return empty object and error.
		return &response.GetUserAdressResponse{}, err
	}
	result := &response.GetUserAdressResponse{
		FullName: userAddress.FullName, StreetAddress: userAddress.StreetAddress,
		Country: userAddress.StreetAddress, State: userAddress.State, Zip: userAddress.Zip,
	}
	return result, nil
}

// Createuser method for creating UserAddress by given UserAddress object.
func (q *UserQueries) CreateUserAdress(address *schema.UserAddress) error {
	// Define query string.
	query := `INSERT INTO address (userid,fullName,streetAddress,country,state,zip)
	 VALUES (:userid,:fullName,:streetAddress,:country,:state,:zip) 
	 ON DUPLICATE KEY UPDATE fullName=:fullName,streetAddress=:streetAddress,country=:country,state=:state,zip=:zip`

	// Send query to database.
	_, err := q.NamedExec(query, address)
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
	_, err := q.Exec(query, user.UserName, user.Telephone, user.AuditInfo, userid)
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
