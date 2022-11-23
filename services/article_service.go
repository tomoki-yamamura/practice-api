package services

import (
	"github.com/tomoki-yamamura/practice-api/models"
	"github.com/tomoki-yamamura/practice-api/repositories"
)

func GetArticleService(articleID int) (models.Article, error) {
	// TODO: sql.DB型を手に入れて、変数に格納
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func GetArticleListService(page int) ([]models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	articles, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func PostArticleService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	newArticle, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}
	return newArticle, nil
}

func PostNiceService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()
	err = repositories.UpdateNiceNum(db, article.ID)
	if err != nil {
		return models.Article{}, err
	}
	updateArticle := models.Article{
		ID: article.ID,
		Title: article.Title,
		Contents: article.Contents,
		UserName: article.UserName,
		NiceNum: article.NiceNum + 1,
		CommentList: article.CommentList,
		CreatedAt: article.CreatedAt,
	}
	return updateArticle, nil
}

// func PostNiceService(article *models.Article) (models.Article, error) {
// 	db, err := connectDB()
// 	if err != nil {
// 		return models.Article{}, err
// 	}
// 	error := repositories.UpdateNiceNum(db, article.ID)
// 	if error != nil {
// 		return models.Article{}, error
// 	}
// 	article.NiceNum += 1
// 	return *article, nil
// }