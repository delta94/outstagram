package cmtablecontroller

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"outstagram/server/dtos/cmtabledtos"
	postVisibility "outstagram/server/enums/postvisibility"
	"outstagram/server/models"
	"outstagram/server/services/cmtableservice"
	"outstagram/server/services/cmtservice"
	"outstagram/server/services/rctableservice"
	"outstagram/server/services/replyservice"
	"outstagram/server/services/userservice"
	"outstagram/server/utils"
)

type Controller struct {
	commentableService *cmtableservice.CommentableService
	commentService     *cmtservice.CommentService
	userService        *userservice.UserService
	reactableService   *rctableservice.ReactableService
	replyService       *replyservice.ReplyService
}

func New(commentableService *cmtableservice.CommentableService, commentService *cmtservice.CommentService, userService *userservice.UserService, reactableService *rctableservice.ReactableService, replyService *replyservice.ReplyService) *Controller {
	return &Controller{commentableService: commentableService, commentService: commentService, userService: userService, reactableService: reactableService, replyService: replyService}
}

//getDTOComment maps comment into a DTO object
func (cc *Controller) getDTOComment(comment *models.Comment) cmtabledtos.Comment {
	return cmtabledtos.Comment{
		ID:            comment.ID,
		Content:       comment.Content,
		ReplyCount:    comment.ReplyCount,
		CreatedAt:     comment.CreatedAt,
		OwnerFullname: comment.User.Fullname,
		OwnerID:       comment.UserID,
		ReactCount:    cc.reactableService.GetReactCount(comment.ReactableID),
		Reactors:      cc.reactableService.GetReactors(comment.ReactableID)}
}

// getDTOReply maps a reply into a DTO object
func (cc *Controller) getDTOReply(reply *models.Reply) cmtabledtos.Reply {
	return cmtabledtos.Reply{
		ID:            reply.ID,
		Content:       reply.Content,
		CreatedAt:     reply.CreatedAt,
		OwnerID:       reply.UserID,
		OwnerFullname: reply.User.Fullname}
}

// checkUserAuthorizationForCommentable checks if user has the authorization to see the post
func (cc *Controller) checkUserAuthorizationForCommentable(cmtableID, userID uint) *utils.HttpError {
	visibility, ownerID, err := cc.commentableService.GetVisibilityByID(cmtableID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return utils.NewHttpError(http.StatusNotFound, "Commentable not found", err.Error())
		}

		return utils.NewHttpError(http.StatusInternalServerError, "Error while retrieving post", err.Error())
	}

	if visibility == postVisibility.Private {
		if ownerID != userID {
			return utils.NewHttpError(http.StatusForbidden, "Cannot access this commentable", nil)
		}
	} else if visibility == postVisibility.OnlyFollowers {
		followed, err := cc.userService.CheckFollow(userID, ownerID)
		if err != nil {
			return utils.NewHttpError(http.StatusInternalServerError, "Error while checking follow", err.Error())
		}

		if !followed {
			return utils.NewHttpError(http.StatusForbidden, "Cannot access this commentable", nil)

		}
	} else if visibility != postVisibility.Public {
		return utils.NewHttpError(http.StatusConflict, "Invalid visibility of a commentable", visibility)
	}

	return nil
}
