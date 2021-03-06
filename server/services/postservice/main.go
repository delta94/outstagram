package postservice

import (
	"github.com/jinzhu/gorm"
	"outstagram/server/constants"
	"outstagram/server/dtos/dtomodels"
	"outstagram/server/enums/postprivacy"
	"outstagram/server/models"
	"outstagram/server/repos/postrepo"
	"outstagram/server/services/cmtableservice"
	"outstagram/server/services/rctableservice"
	"outstagram/server/services/userservice"
	"outstagram/server/utils"
)

type PostService struct {
	postRepo           *postrepo.PostRepo
	userService        *userservice.UserService
	reactableService   *rctableservice.ReactableService
	commentableService *cmtableservice.CommentableService
}

func New(postRepo *postrepo.PostRepo, userService *userservice.UserService, reactableService *rctableservice.ReactableService, commentableService *cmtableservice.CommentableService) *PostService {
	return &PostService{postRepo: postRepo, userService: userService, reactableService: reactableService, commentableService: commentableService}
}

func (s *PostService) Save(post *models.Post) error {
	return s.postRepo.Save(post)
}

func (s *PostService) Update(post *models.Post, values map[string]interface{}) error {
	return s.postRepo.Update(post, values)
}

// GetUserPosts return array of all posts that user has
func (s *PostService) GetUserPosts(userID uint) ([]models.Post, error) {
	posts, err := s.postRepo.GetPostsByUserID(userID)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// GetUsersPostsWithLimit returns array of posts with their basic info
func (s *PostService) GetUsersPostsWithLimit(userID, limit, offset uint) ([]models.Post, error) {
	posts, err := s.postRepo.GetPostsByUserIDWithLimit(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// GetPostByID lets user get the post that has the postID specified in parameter
// User may be restricted to view the post due to its visibility. In such case, ErrRecordNotFound is returned.
// `userID` is the id of user who wants to view the post
func (s *PostService) GetPostByID(postID, audienceID uint) (*models.Post, error) {
	post, err := s.postRepo.FindByID(postID)

	if err != nil {
		return nil, err
	}

	if audienceID == post.UserID {
		return post, nil
	}

	if post.Privacy == postPrivacy.Public {
		return post, nil
	}

	if post.Privacy == postPrivacy.Private {
		return nil, gorm.ErrRecordNotFound
	}

	if audienceID == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	// If post.Privacy is OnlyFollowers
	ok, err := s.userService.CheckFollow(audienceID, post.UserID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, gorm.ErrRecordNotFound
	}

	return post, nil
}

func (s *PostService) Search(text string) ([]*models.Post, error) {
	return s.postRepo.Search(text)
}

// getDTOPost maps post, including post's images, post's comments into a DTO object
func (s *PostService) GetDTOPost(post *models.Post, userID, audienceUserID uint) (*dtomodels.Post, error) {
	// Set basic post's info

	var imageCount int
	if len(post.Images) > 0 {
		imageCount = len(post.Images)
	} else {
		imageCount = 1
	}

	dtoPost := dtomodels.Post{
		ID:            post.ID,
		ViewableID:    post.ViewableID,
		CommentableID: post.CommentableID,
		ReactableID:   post.ReactableID,
		CreatedAt:     post.CreatedAt,
		Content:       post.Content,
		Visibility:    post.Privacy,
		ImageCount:    imageCount,
		OwnerID:       post.UserID,
		OwnerFullname: post.User.Fullname,
		OwnerUsername: post.User.Username,
		ReactCount:    s.reactableService.GetReactCount(post.ReactableID),
		Reacted:       s.reactableService.CheckUserReaction(audienceUserID, post.ReactableID),
		Reactors:      s.reactableService.GetReactorDTOs(post.ReactableID, audienceUserID, 5, 0),
	}

	if imageCount == 1 {
		dtoPost.ImageID = utils.NewUintPointer(post.Image.ID)
	} else {
		for _, postImage := range post.Images {
			dtoPost.Images = append(dtoPost.Images, dtomodels.SimplePostImage{ID: postImage.ID, ImageID: postImage.ImageID})
		}
	}

	commentable, err := s.commentableService.GetCommentsWithLimit(post.CommentableID, userID, constants.PostCommentCount, 0)
	if err != nil {
		return nil, err
	}

	dtoPost.CommentCount = commentable.CommentCount
	for _, comment := range commentable.Comments {
		dtoComment := comment.ToDTO()
		dtoComment.Reacted = s.reactableService.CheckUserReaction(audienceUserID, comment.ReactableID)
		dtoPost.Comments = append(dtoPost.Comments, dtoComment)
	}

	return &dtoPost, nil
}
