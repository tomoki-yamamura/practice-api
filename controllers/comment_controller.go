package controllers

import "github.com/tomoki-yamamura/practice-api/controllers/services"

type CommentController struct {
	services services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{services: s}
}
