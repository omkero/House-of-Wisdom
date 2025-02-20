package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"wisdom/src/config"
	"wisdom/src/models"
	"wisdom/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type ArticleRepository struct{}

func (arp *ArticleRepository) Select_all_Articles_repo(limit int, offset int) ([]models.Article, error) {
	// Check if the pool is initialized
	if config.Pool == nil {
		log.Fatal("Database connection pool is not initialized")
		return nil, nil
	}

	var query string = `SELECT Article_id,Article_title,Article_content,Created_at,Banner_img,Tags,Views,User_id,Username,Updated_at,public_article_title, description FROM articles WHERE is_hidden = 'false'  ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	// Query the database using the connection pool
	rows, err := config.Pool.Query(context.Background(), query, limit, offset)
	if err != nil {
		fmt.Printf("Error querying articles: %v", err)
		return nil, err
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var article_body models.Article

		err = rows.Scan(
			&article_body.Article_id,
			&article_body.Article_title,
			&article_body.Article_content,
			&article_body.Created_at,
			&article_body.Banner_img,
			&article_body.Tags,
			&article_body.Views,
			//	&article_body.Is_hidden,
			&article_body.User_id,
			&article_body.Username,
			&article_body.Updated_at,
			&article_body.Public_article_title,
			&article_body.Description,
		)

		if err != nil {
			fmt.Printf("Error scanning row: %v", err)
			return nil, err
		}
		articles = append(articles, article_body)
	}

	return articles, nil
}

func (arp *ArticleRepository) Select_Article_repo(article_title string) (models.Article, error) {
	// Check if the pool is initialized
	if config.Pool == nil {
		log.Fatal("Database connection pool is not initialized")
		return models.Article{}, nil
	}

	var article models.Article
	var query string = `SELECT Article_id,Article_title,Article_content,Created_at,Banner_img,Tags,Views,User_id,Username,Updated_at,public_article_title, description FROM articles WHERE public_article_title = $1 and is_hidden = 'false'`

	// Query the database using the connection pool
	err := config.Pool.QueryRow(context.Background(), query, article_title).Scan(&article.Article_id, &article.Article_title, &article.Article_content, &article.Created_at, &article.Banner_img, &article.Tags, &article.Views, &article.User_id, &article.Username, &article.Updated_at, &article.Public_article_title, &article.Description)
	if err != nil {
		fmt.Printf("Error querying articles: %v", err)
		return models.Article{}, err
	}

	return article, nil
}

func (arp *ArticleRepository) Insert_new_article(ArticleBody models.CreateArticle, username string, user_id int, path string) (string, error) {

	if config.Pool == nil {
		log.Fatal("Error cannot connect to database")
		return "cannot establish connection", nil
	}

	var public_article_title string = strings.Replace(ArticleBody.Article_title, " ", "-", -1)

	_, err := config.Pool.Exec(context.Background(), "INSERT INTO articles (article_title, article_content, banner_img, tags,  user_id, username, public_article_title, description) VALUES ($1,$2,$3, $4, $5, $6, $7, $8)", ArticleBody.Article_title, ArticleBody.Article_content, path, ArticleBody.Tags, user_id, username, public_article_title, ArticleBody.Description)

	if err != nil {
		return "error has occurred: " + err.Error(), err
	}

	return "article has been created", nil
}

func (arp *ArticleRepository) Delete_article(artilce_id models.ArticleID, username string) (string, int, error) {
	if config.Pool == nil {
		log.Fatal("Error cannot connect to database")
		return "cannot establish connection", 500, nil
	}

	var articleID models.ArticleID

	row := config.Pool.QueryRow(context.Background(), "SELECT article_id, user_id, banner_img, username FROM articles WHERE article_id = $1", artilce_id.Article_id).Scan(&articleID.Article_id, &articleID.User_id, &articleID.Banner_img, &articleID.Username)
	if row == pgx.ErrNoRows {
		return "", 404, errors.New("there is no article found with this article_id")
	}

	if username != articleID.Username {
		return "", 401, errors.New("error you cannot delete an article that not created by you")
	}

	// remove the file from the disk
	var dirPath = "public/static/" + articleID.Banner_img
	errr := os.Remove(dirPath)
	if errr != nil {
		fmt.Println(errr)
	}

	_, err := config.Pool.Exec(context.Background(), "DELETE FROM articles WHERE article_id = $1 and username = $2", artilce_id.Article_id, username)
	if err != nil {
		return "", 500, err
	}

	return "Article has been deleted", 201, nil
}

func (arp *ArticleRepository) Update_article(artilce models.EditArticleForm, user_id int, ctx *gin.Context) (string, int, error) {
	if config.Pool == nil {
		log.Fatal("Error cannot connect to database")
		return "cannot establish connection", 500, nil
	}

	var banner_img models.EditArticleForm

	row := config.Pool.QueryRow(context.Background(), "SELECT banner_img FROM articles WHERE article_id = $1", artilce.Article_id).Scan(&banner_img.Banner_img)

	if row == pgx.ErrNoRows {
		return "", 404, errors.New("there is no article found with this article_id")
	}
	var dirPath = "public/static/" + banner_img.Banner_img
	err := os.Remove(dirPath)
	if err != nil {
		fmt.Println(err)
	}

	publicPath, _, _, errB := utils.SaveImageToDirEdit(&artilce, ctx)
	if errB != nil {
		fmt.Println(err)
	}

	var public_article_title string = strings.Replace(artilce.Article_title, " ", "-", -1)

	res, err := config.Pool.Exec(context.Background(), "UPDATE articles SET article_title = $1, article_content = $2, banner_img = $3, description = $4, public_article_title = $5, tags = $6, updated_at = NOW()  WHERE article_id = $7 and user_id = $8", artilce.Article_title, artilce.Article_content, publicPath, artilce.Description, public_article_title, artilce.Tags, artilce.Article_id, user_id)

	if res.String() == "UPDATE 0" {
		return "", 401, errors.New("error you cannot edit an article that not created by you")
	}
	if err != nil {
		return "", 500, err
	}

	return "Article has been updated", 201, nil
}
