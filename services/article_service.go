package services

import (
	"fmt"

	"github.com/tomoki-yamamura/practice-api/models"
	"github.com/tomoki-yamamura/practice-api/repositories"
)

// PostArticleHandlerで使うことを想定したサービス
// 引数の情報をもとに新しい記事を作り、結果を返却
func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		return models.Article{}, err
	}
	return newArticle, nil
}

// ArticleListHandlerで使うことを想定したサービス
// 指定pageの記事一覧を返却
func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	fmt.Println("GetArticleListService", s)

	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		return nil, err
	}

	return articleList, nil
}

// ArticleDetailHandlerで使うことを想定したサービス
// 指定IDの記事情報を返却
func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}
	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

// PostNiceHandlerで使うことを想定したサービス
// 指定IDの記事のいいね数を+1して、結果を返却
// func PostNiceService(article models.Article) (models.Article, error) {
// 	db, err := connectDB()
// 	if err != nil {
// 		return models.Article{}, err
// 	}
// 	defer db.Close()

// 	err = repositories.UpdateNiceNum(db, article.ID)
// 	if err != nil {
// 		return models.Article{}, err
// 	}

// 	return models.Article{
// 		ID:        article.ID,
// 		Title:     article.Title,
// 		Contents:  article.Contents,
// 		UserName:  article.UserName,
// 		NiceNum:   article.NiceNum + 1,
// 		CreatedAt: article.CreatedAt,
// 	}, nil
// }

func (s *MyAppService) PostNiceService(article *models.Article) (models.Article, error) {
	fmt.Println("PostNiceService", s.db)
	error := repositories.UpdateNiceNum(s.db, article.ID)
	if error != nil {
		return models.Article{}, error
	}
	article.UpNiceNum()
	return *article, nil
}
