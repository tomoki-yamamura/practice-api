package controllers

import (
	"github.com/tomoki-yamamura/practice-api/controllers/services"
)

type ArticleController struct {
	service services.ArticleServicer
}

func NewArticleController(s services.ArticleServicer) *ArticleController {
	return &ArticleController{service: s}
}