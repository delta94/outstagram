// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package configs

import (
	"outstagram/server/controllers/authcontroller"
	"outstagram/server/controllers/cmtablecontroller"
	"outstagram/server/controllers/postcontroller"
	"outstagram/server/controllers/rctcontroller"
	"outstagram/server/controllers/usercontroller"
	"outstagram/server/db"
	"outstagram/server/repos/cmtablerepo"
	"outstagram/server/repos/cmtrepo"
	"outstagram/server/repos/imgrepo"
	"outstagram/server/repos/notifbrepo"
	"outstagram/server/repos/postimgrepo"
	"outstagram/server/repos/postrepo"
	"outstagram/server/repos/rctablerepo"
	"outstagram/server/repos/rctrepo"
	"outstagram/server/repos/replyrepo"
	"outstagram/server/repos/storybrepo"
	"outstagram/server/repos/userrepo"
	"outstagram/server/repos/vwablerepo"
	"outstagram/server/services/cmtableservice"
	"outstagram/server/services/cmtservice"
	"outstagram/server/services/imgservice"
	"outstagram/server/services/notifbservice"
	"outstagram/server/services/postimgservice"
	"outstagram/server/services/postservice"
	"outstagram/server/services/rctableservice"
	"outstagram/server/services/rctservice"
	"outstagram/server/services/replyservice"
	"outstagram/server/services/storybservice"
	"outstagram/server/services/userservice"
	"outstagram/server/services/vwableservice"
)

// Injectors from injection.go:

func InitializeUserController() (*usercontroller.Controller, error) {
	gormDB, err := db.New()
	if err != nil {
		return nil, err
	}
	userRepo := userrepo.New(gormDB)
	userService := userservice.New(userRepo)
	controller := usercontroller.New(userService)
	return controller, nil
}

func InitializeAuthController() (*authcontroller.Controller, error) {
	gormDB, err := db.New()
	if err != nil {
		return nil, err
	}
	userRepo := userrepo.New(gormDB)
	userService := userservice.New(userRepo)
	notifBoardRepo := notifbrepo.New(gormDB)
	notifBoardService := notifbservice.New(notifBoardRepo)
	storyBoardRepo := storybrepo.New(gormDB)
	storyBoardService := storybservice.New(storyBoardRepo)
	controller := authcontroller.New(userService, notifBoardService, storyBoardService)
	return controller, nil
}

func InitializePostController() (*postcontroller.Controller, error) {
	gormDB, err := db.New()
	if err != nil {
		return nil, err
	}
	postRepo := postrepo.New(gormDB)
	userRepo := userrepo.New(gormDB)
	userService := userservice.New(userRepo)
	postService := postservice.New(postRepo, userService)
	imageRepo := imgrepo.New(gormDB)
	imageService := imgservice.New(imageRepo)
	postImageRepo := postimgrepo.New(gormDB)
	postImageService := postimgservice.New(postImageRepo)
	commentableRepo := cmtablerepo.New(gormDB)
	commentableService := cmtableservice.New(commentableRepo)
	commentRepo := cmtrepo.New(gormDB)
	commentService := cmtservice.New(commentRepo)
	reactableRepo := rctablerepo.New(gormDB)
	reactableService := rctableservice.New(reactableRepo)
	viewableRepo := vwablerepo.New(gormDB)
	viewableService := vwableservice.New(viewableRepo)
	replyRepo := replyrepo.New(gormDB)
	replyService := replyservice.New(replyRepo)
	controller := postcontroller.New(postService, imageService, postImageService, commentableService, commentService, reactableService, userService, viewableService, replyService)
	return controller, nil
}

func InitializeReactController() (*rctcontroller.Controller, error) {
	gormDB, err := db.New()
	if err != nil {
		return nil, err
	}
	reactRepo := rctrepo.New(gormDB)
	reactService := rctservice.New(reactRepo)
	controller := rctcontroller.New(reactService)
	return controller, nil
}

func InitializeCommentableController() (*cmtablecontroller.Controller, error) {
	gormDB, err := db.New()
	if err != nil {
		return nil, err
	}
	commentableRepo := cmtablerepo.New(gormDB)
	commentableService := cmtableservice.New(commentableRepo)
	commentRepo := cmtrepo.New(gormDB)
	commentService := cmtservice.New(commentRepo)
	userRepo := userrepo.New(gormDB)
	userService := userservice.New(userRepo)
	reactableRepo := rctablerepo.New(gormDB)
	reactableService := rctableservice.New(reactableRepo)
	replyRepo := replyrepo.New(gormDB)
	replyService := replyservice.New(replyRepo)
	controller := cmtablecontroller.New(commentableService, commentService, userService, reactableService, replyService)
	return controller, nil
}
