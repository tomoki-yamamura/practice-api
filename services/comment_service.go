package services

import (
	"github.com/tomoki-yamamura/practice-api/apperrors"
	"github.com/tomoki-yamamura/practice-api/models"
	"github.com/tomoki-yamamura/practice-api/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDetailFailed.Wrap(err, "failed insert")
		return models.Comment{}, err
	}
	return newComment, nil
}
