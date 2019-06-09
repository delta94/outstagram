// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package configs

import (
	"outstagram/server/controllers/authcontroller"
	"outstagram/server/controllers/postcontroller"
	"outstagram/server/controllers/usercontroller"
	"outstagram/server/db"
	"outstagram/server/repos/cmtablerepo"
	"outstagram/server/repos/cmtrepo"
	"outstagram/server/repos/imgrepo"
	"outstagram/server/repos/notifbrepo"
	"outstagram/server/repos/postimgrepo"
	"outstagram/server/repos/postrepo"
	"outstagram/server/repos/rctablerepo"
	"outstagram/server/repos/storybrepo"
	"outstagram/server/repos/userrepo"
	"outstagram/server/services/cmtableservice"
	"outstagram/server/services/cmtservice"
	"outstagram/server/services/imgservice"
	"outstagram/server/services/notifbservice"
	"outstagram/server/services/postimgservice"
	"outstagram/server/services/postservice"
	"outstagram/server/services/rctableservice"
	"outstagram/server/services/storybservice"
	"outstagram/server/services/userservice"
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
	postService := postservice.New(postRepo)
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
	controller := postcontroller.New(postService, imageService, postImageService, commentableService, commentService, reactableService)
	return controller, nil
}
