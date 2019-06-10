//+build wireinject

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
	"outstagram/server/repos/viewablerepo"
	"outstagram/server/services/cmtableservice"
	"outstagram/server/services/cmtservice"
	"outstagram/server/services/imgservice"
	"outstagram/server/services/notifbservice"
	"outstagram/server/services/postimgservice"
	"outstagram/server/services/postservice"
	"outstagram/server/services/rctableservice"
	"outstagram/server/services/storybservice"
	"outstagram/server/services/userservice"
	"outstagram/server/services/viewableservice"

	"github.com/google/wire"
)

func InitializeUserController() (*usercontroller.Controller, error) {
	wire.Build(
		usercontroller.New,

		userservice.New,
		userrepo.New,

		db.New)
	return &usercontroller.Controller{}, nil
}

func InitializeAuthController() (*authcontroller.Controller, error) {
	wire.Build(
		authcontroller.New,

		userservice.New,
		userrepo.New,

		notifbservice.New,
		notifbrepo.New,

		storybservice.New,
		storybrepo.New,

		db.New)
	return &authcontroller.Controller{}, nil
}

func InitializePostController() (*postcontroller.Controller, error) {
	wire.Build(
		postcontroller.New,

		viewablerepo.New,
		viewableservice.New,

		userservice.New,
		userrepo.New,

		postservice.New,
		postrepo.New,

		postimgservice.New,
		postimgrepo.New,

		imgservice.New,
		imgrepo.New,

		cmtableservice.New,
		cmtablerepo.New,

		cmtservice.New,
		cmtrepo.New,

		rctableservice.New,
		rctablerepo.New,

		db.New)
	return &postcontroller.Controller{}, nil
}
