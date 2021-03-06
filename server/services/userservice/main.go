package userservice

import (
	"errors"
	"github.com/jinzhu/gorm"
	"net/http"
	"outstagram/server/constants"
	"outstagram/server/models"
	"outstagram/server/repos/userrepo"
	"outstagram/server/utils"
)

type UserService struct {
	userRepo *userrepo.UserRepo
}

func New(userRepo *userrepo.UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) FindByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *UserService) FindByUsername(username string) (*models.User, error) {
	return s.userRepo.FindByUsername(username)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}

func (s *UserService) VerifyLogin(username, password string) (*models.User, *models.AppError) {
	user, err := s.userRepo.FindByUsername(username)
	if gorm.IsRecordNotFoundError(err) {
		id := "username_not_found"
		where := "UserService.VerifyLogin"
		message := "Username not found"
		params := map[string]interface{}{"username": username}
		code := http.StatusNotFound
		return nil, models.NewAppError(id, where, message, params, code)
	}

	if user.Password != password {
		id := "password_incorrect"
		where := "UserService.VerifyLogin"
		message := "Password is incorrect"
		params := map[string]interface{}{"password": password}
		code := http.StatusConflict
		return nil, models.NewAppError(id, where, message, params, code)
	}

	return user, nil
}

func (s *UserService) Search(text string, options ...map[string]interface{}) ([]*models.User, error) {
	return s.userRepo.Search(text, options...)
}

func (s *UserService) Save(user *models.User) error {
	if s.CheckExistsByID(user.ID) {
		return s.userRepo.Save(user)
	}

	return s.userRepo.Create(user)
}

func (s *UserService) Delete(id uint) error {
	return s.userRepo.DeleteByID(id)
}

func (s *UserService) CheckExistsByID(id uint) bool {
	return s.userRepo.ExistsByID(id)
}

func (s *UserService) CheckExistsByUsername(username string) bool {
	return s.userRepo.ExistsByUsername(username)
}

func (s *UserService) CheckExistsByEmail(email string) bool {
	return s.userRepo.ExistsByEmail(email)
}

func (s *UserService) GetFollowers(userID uint) []models.User {
	return s.userRepo.GetFollowers(userID)
}

func (s *UserService) GetFollowings(userID uint) []*models.User {
	return s.userRepo.GetFollowings(userID)
}

func (s *UserService) CheckFollow(follow, followed uint) (bool, error) {
	return s.userRepo.CheckFollow(follow, followed)
}

func (s *UserService) Follow(following, follower uint) error {
	hasFollowed, err := s.CheckFollow(following, follower)
	if err != nil {
		return err
	}

	if hasFollowed {
		return errors.New(constants.AlreadyExist)
	}

	return s.userRepo.Follow(following, follower)
}

func (s *UserService) Unfollow(following, follower uint) error {
	hasFollowed, err := s.CheckFollow(following, follower)
	if err != nil {
		return err
	}

	if !hasFollowed {
		return errors.New(constants.NotExist)
	}

	return s.userRepo.Unfollow(following, follower)
}

func (s *UserService) GetPostFeed(userID uint) []models.Post {
	return s.userRepo.GetPostFeed(userID)
}

func (s *UserService) GetFollowingsWithAffinity(userID uint) []*models.User {
	return s.userRepo.GetFollowingsWithAffinity(userID)
}

func (s *UserService) GetUserByUserIDOrUsername(userIDOrUsername interface{}) (*models.User, error) {
	var user *models.User
	var err error

	if username, ok := userIDOrUsername.(string); ok {
		id, err := utils.StringToUint(username)
		if err != nil {
			user, err = s.FindByUsername(username)
		} else {
			user, err = s.FindByID(id)
		}
	}

	if id, ok := userIDOrUsername.(uint); ok {
		user, err = s.FindByID(id)
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserRoomIDs(userID uint) ([]uint, error) {
	return s.userRepo.GetUserRoomIDs(userID)
}

func (s *UserService) GetFollowSuggestions(userID uint) []*models.User {
	return s.userRepo.GetFollowSuggestions(userID)
}