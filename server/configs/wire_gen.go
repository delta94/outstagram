// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package configs

import (
	"outstagram/server/controllers/authcontroller"
	"outstagram/server/controllers/usercontroller"
	"outstagram/server/db"
	"outstagram/server/repositories/nbrepo"
	"outstagram/server/repositories/sbrepo"
	"outstagram/server/repositories/userrepo"
	"outstagram/server/services/nbservice"
	"outstagram/server/services/sbservice"
	"outstagram/server/services/userservice"
)

// Injectors from injection.go:

func InitializeUserController() (*usercontroller.Controller, error) {
	gormDB, err := db.New()
	if err != nil {
		return nil, err
	}
	userRepository := userrepo.New(gormDB)
	userService := userservice.New(userRepository)
	controller := usercontroller.New(userService)
	return controller, nil
}

func InitializeAuthController() (*authcontroller.Controller, error) {
	gormDB, err := db.New()
	if err != nil {
		return nil, err
	}
	userRepository := userrepo.New(gormDB)
	userService := userservice.New(userRepository)
	notifBoardRepo := nbrepo.New(gormDB)
	notifBoardService := nbservice.New(notifBoardRepo)
	storyBoardRepo := sbrepo.New(gormDB)
	storyBoardService := sbservice.New(storyBoardRepo)
	controller := authcontroller.New(userService, notifBoardService, storyBoardService)
	return controller, nil
}