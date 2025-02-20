package repositories

import (
	"context"
	"errors"
	"fmt"
	"wisdom/src/config"
	"wisdom/src/models"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepository struct{}

func (atp *AuthRepository) Select_user_data(data models.AuthLoginModel) (models.AuthUserData, error) {
	// Check database connection
	if config.Pool == nil {
		return models.AuthUserData{}, fmt.Errorf("error: cannot establish database connection")
	}

	var UserAuth models.AuthUserData

	// Use explicit column names
	query := `SELECT * FROM users WHERE email_address = $1`

	// Execute the query
	err := config.Pool.QueryRow(context.Background(), query, data.Email).Scan(
		&UserAuth.User_id,
		&UserAuth.Username,
		&UserAuth.Email_address,
		&UserAuth.Password,
		&UserAuth.Is_banned,
		&UserAuth.Is_suspended,
		&UserAuth.Articles_count,
		&UserAuth.Profile_img,
		&UserAuth.Created_at,
	)

	if err != nil {
		// Log and return the error
		return models.AuthUserData{}, err
	}

	// Return the user data
	return UserAuth, nil
}

func (atp *AuthRepository) Verify_user_exist_by_email(data models.AuthSignupModel) (bool, error) {
	if config.Pool == nil {
		return false, errors.New("error cannot connect to pool")
	}

	var existingEmail, existingUsername string

	query := `SELECT email_address, username 
              FROM users 
              WHERE email_address = $1 OR username = $2`
	err := config.Pool.QueryRow(context.Background(), query, data.Email, data.Username).Scan(&existingEmail, &existingUsername)

	if err != nil {
		if err == pgx.ErrNoRows {
			return false, nil // No user found
		}
		return false, err // Database error
	}

	// Return true if either email or username exists
	if existingEmail == data.Email || existingUsername == data.Username {
		return true, nil
	}

	return false, nil
}

func (atp *AuthRepository) Insert_new_user(data models.AuthSignupModel) bool {
	if config.Pool == nil {
		return false
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return false
	}

	_, err = config.Pool.Exec(context.Background(), "INSERT INTO users (email_address, username, password) VALUES($1, $2, $3)", data.Email, data.Username, string(hashedPassword))

	if err != nil {
		return false
	}

	return true
}
