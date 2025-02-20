package models

import (
	"database/sql"
	"time"
)

type AuthLoginModel struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthSignupModel struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthEmail struct {
	Email    string `json:"email_address" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type AuthUserData struct {
	User_id        int           `json:"user_id"`
	Username       string        `json:"username"`
	Email_address  string        `json:"email_address"`
	Password       string        `json:"password"`
	Is_banned      bool          `json:"is_banned"`
	Is_suspended   bool          `json:"is_suspended"`
	Articles_count sql.NullInt64 `json:"articles_count"`
	Profile_img    string        `json:"profile_img"`
	Created_at     time.Time     `json:"created_at"`
}
