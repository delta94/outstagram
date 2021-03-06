package postcontroller

import (
	"outstagram/server/services/imgservice"
	"outstagram/server/services/postimgservice"
	"outstagram/server/services/postservice"
	"outstagram/server/services/userservice"
	"outstagram/server/services/vwableservice"
)

type Controller struct {
	postService      *postservice.PostService
	imageService     *imgservice.ImageService
	postImageService *postimgservice.PostImageService
	viewableService  *vwableservice.ViewableService
	userService      *userservice.UserService
}

func New(postService *postservice.PostService,
	imageService *imgservice.ImageService,
	postImageService *postimgservice.PostImageService,
	viewableService *vwableservice.ViewableService,
	userService *userservice.UserService) *Controller {
	return &Controller{
		postService:      postService,
		imageService:     imageService,
		postImageService: postImageService,
		viewableService:  viewableService,
		userService:      userService,
	}
}
