package models

import (
	"mime/multipart"
	"time"
)

type Article struct {
	Article_id      int      `json:"article_id"`
	Article_title   string   `json:"article_title"`
	Article_content string   `json:"article_content"`
	Banner_img      string   `json:"banner_img"` // use sql.NullString from "database/sql" if you expected null from the database or
	Tags            []string `json:"tags"`       // you will get error
	Views           int      `json:"views"`
	// Is_hidden       bool           `json:"is_hidden"`
	User_id              int       `json:"user_id"`
	Username             string    `json:"username"`
	Created_at           time.Time `json:"created_at"`
	Updated_at           time.Time `json:"updated_at"`
	Public_article_title string    `json:"public_article_title"`
	Description          string    `json:"description"`
}

type ArticleID struct {
	Article_id int    `json:"article_id"`
	User_id    int    `json:"user_id"`
	Username   string `json:"username"`
	Banner_img string `json:"banner_img"`
}

type EditArticle struct {
	Article_id      int      `json:"article_id"`
	Article_title   string   `json:"article_title"`
	Article_content string   `json:"article_content"`
	Banner_img      string   `json:"banner_img"`
	Tags            []string `json:"tags"`
}

type EditArticleForm struct {
	Article_id      int                   `form:"article_id" binding:"required"`
	Article_title   string                `form:"article_title" binding:"required"`
	Article_content string                `form:"article_content"`
	Banner_img      string                `form:"banner_img" binding:"required"`
	File            *multipart.FileHeader `form:"file" binding:"required"`
	Tags            []string              `form:"tags" binding:"required"`
	Description     string                `form:"description" binding:"required"`
}

type GetArticlesLIMIT struct {
	Page int `json:"page_number"`
}

type GetArticle struct {
	Article_title string `json:"article_title" binding:"required"`
}

type CreateArticle struct {
	Article_title   string                `form:"article_title" binding:"required"`
	Tags            string                `form:"tags" binding:"required"`
	Article_content string                `form:"article_content" binding:"required"`
	Banner_img      string                `form:"banner_img" binding:"required"`
	File            *multipart.FileHeader `form:"file" binding:"required"`
	Description     string                `form:"description" binding:"required"`
}
