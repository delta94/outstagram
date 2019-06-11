package cmtableservice

import (
	postVisibility "outstagram/server/enums/postvisibility"
	"outstagram/server/models"
	"outstagram/server/repos/cmtablerepo"
)

type CommentableService struct {
	commentableRepo *cmtablerepo.CommentableRepo
}

func New(commentableRepo *cmtablerepo.CommentableRepo) *CommentableService {
	return &CommentableService{commentableRepo: commentableRepo}
}

func (s *CommentableService) GetCommentCount(commentableID uint) int {
	return s.commentableRepo.GetCommentCount(commentableID)
}

func (s *CommentableService) GetCommentsWithLimit(id uint, limit uint, offset uint) (*models.Commentable, error) {
	return s.commentableRepo.GetCommentsWithLimit(id, limit, offset)
}

func (s *CommentableService) GetComments(id uint) (*models.Commentable, error) {
	return s.commentableRepo.GetComments(id)
}

func (s *CommentableService) GetVisibilityByID(id uint) (postVisibility.Visibility, uint, error) {
	return s.commentableRepo.GetVisibility(id)
}

func (s *CommentableService) CheckHasComment(cmtableID, cmtID uint) bool {
	return s.commentableRepo.HasComment(cmtableID, cmtID)
}
