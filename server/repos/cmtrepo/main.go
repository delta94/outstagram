package cmtrepo

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/models"
)

type CommentRepo struct {
	db *gorm.DB
}

func New(dbConnection *gorm.DB) *CommentRepo {
	return &CommentRepo{db: dbConnection}
}

func (r *CommentRepo) Save(comment *models.Comment) error {
	reactable := models.Reactable{}
	r.db.Create(&reactable)
	comment.ReactableID = reactable.ID
	err := r.db.Create(&comment).Error
	if err != nil {
		return err
	}

	// WARNING: This is for retrieve the information of the comment's owner
	r.db.Model(&comment).Related(&comment.User)
	return nil
}

func (r *CommentRepo) FindByID(id uint) (*models.Comment, error) {
	var comment models.Comment

	if err := r.db.First(&comment, id).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *CommentRepo) GetReplies(id uint) (*models.Comment, error) {
	var comment models.Comment

	if err := r.db.First(&comment, id).Error; err != nil {
		return nil, err
	}

	r.db.Model(&comment).Related(&comment.Replies)
	replies := comment.Replies
	for i := 0; i < len(replies); i++ {
		r.db.Model(&replies[i]).Related(&replies[i].User)
	}

	return &comment, nil
}

func (r *CommentRepo) GetRepliesWithLimit(id, limit, offset uint) (*models.Comment, error) {
	var comment models.Comment

	if err := r.db.First(&comment, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Raw(`
	SELECT * 
		FROM (SELECT * FROM reply WHERE comment_id = ? ORDER BY created_at DESC LIMIT ? OFFSET ?) AS reversed
	ORDER BY created_at ASC
	`, id, limit, offset).Scan(&comment.Replies).Error; err != nil {
		return nil, err
	}

	replies := comment.Replies
	for i := 0; i < len(replies); i++ {
		r.db.Model(&replies[i]).Related(&replies[i].User)
	}

	comment.ReplyCount = r.GetReplyCount(id)
	return &comment, nil
}

func (r *CommentRepo) GetReplyCount(id uint) int {
	var count int
	r.db.Model(&models.Comment{}).Where("comment_id = ?", id).Count(&count)
	return count
}
