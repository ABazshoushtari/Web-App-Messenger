package user

import (
	"context"
	"errors"
	"github.com/ABazshoushtari/Web-App-Messenger/domain"
	"github.com/ABazshoushtari/Web-App-Messenger/domain/payloads"
	"github.com/ABazshoushtari/Web-App-Messenger/internal/helpers"
	"github.com/ABazshoushtari/Web-App-Messenger/internal/logger"
	"github.com/ABazshoushtari/Web-App-Messenger/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	repos *repository.Repositories
}

func New(repos *repository.Repositories) *User {
	return &User{
		repos: repos,
	}
}

func (u *User) AuthRegister(ctx context.Context, payload payloads.UserRegisterRequest) (*payloads.UserRegisterResponse, error) {
	if err := u.repos.User.CheckExisting(payload.Username, payload.Phone); !errors.Is(err, gorm.ErrRecordNotFound) {
		if err != nil {
			logger.Logger().Errorw("error while checking if user exists", "error", err)
		}
		return nil, errors.New("a user with same username/phoneNumber exists")
	}

	if len(payload.Password) < 8 {
		return nil, errors.New("password must be at least 8 characters")
	}
	password, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := domain.User{
		Username:    payload.Username,
		Password:    string(password),
		PhoneNumber: payload.Phone,
		FirstName:   payload.FirstName,
		LastName:    payload.LastName,
		Bio:         payload.Bio,
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := u.repos.User.Register(&user); err != nil {
		logger.Logger().Errorw("error while registering user", "error", err)
		return nil, errors.New("error while registering user")
	}

	if err := u.repos.User.SetImage(&user, payload.Image); err != nil {
		logger.Logger().Errorw("error while setting user image", "error", err)
		return nil, errors.New("error while setting user image")
	}

	userDTO := user.ToDTO()
	return &payloads.UserRegisterResponse{User: *userDTO}, nil

}

func (u *User) AuthLogin(ctx context.Context, payload payloads.UserLoginRequest) (*payloads.UserLoginResponse, error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	user := domain.User{}
	if err := u.repos.User.GetByUsername(payload.Username, &user); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid Username")
		}
		logger.Logger().Errorw("error while getting user from db", "error", err)
		return nil, errors.New("error while getting user From DB")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), password); err != nil {
		return nil, errors.New("invalid password")
	}
	token, err := helpers.GenerateJWT(user.ID)
	if err != nil {
		logger.Logger().Errorw("error while generating jwt", "error", err)
		return nil, errors.New("error while generating jwt token")
	}
	return &payloads.UserLoginResponse{
		Token: token,
	}, nil
}

func (u *User) ShowUser(ctx context.Context, userID uint64) (*payloads.UserShowResponse, error) {
	user := domain.User{}
	if err := u.repos.User.GetByID(userID, &user); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		logger.Logger().Errorw("error while getting user from db", "error", err)
		return nil, errors.New("error while getting user from db")
	}
	return &payloads.UserShowResponse{
		UserDTO: user.ToDTO(),
	}, nil
}

func (u *User) IndexUser(ctx context.Context, keyword string) (*payloads.UserShowResponse, error) {
	user := domain.User{}
	if err := u.repos.User.GetByKey(keyword, &user); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		logger.Logger().Errorw("error while getting user from db", "error", err)
		return nil, errors.New("error while getting user from db")
	}
	return &payloads.UserShowResponse{
		UserDTO: user.ToDTO(),
	}, nil

}

func (u *User) UpdateUser(ctx context.Context, userID uint64, payload payloads.UserUpdateRequest) (*payloads.GenericsSuccessFlagResponse, error) {
	user := domain.GetUserDTO(ctx)
	if user == nil {
		return nil, errors.New("invalid jwt user")
	}
	if user.ID != userID {
		return nil, errors.New("input user id does not match with your id")
	}
	updatedUser := user.User
	if payload.Username != "" {
		updatedUser.Username = payload.Username
	}
	if payload.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, errors.New("error while encrypting password")
		}
		updatedUser.Password = string(password)
	}
	if payload.FirstName != "" {
		updatedUser.FirstName = payload.FirstName
	}
	if payload.LastName != "" {
		updatedUser.LastName = payload.LastName
	}
	if payload.Phone != "" {
		updatedUser.PhoneNumber = payload.Phone
	}
	if payload.Bio != "" {
		updatedUser.Bio = payload.Bio
	}
	if err := updatedUser.Validate(); err != nil {
		return nil, err
	}
	if err := u.repos.User.Update(&updatedUser); err != nil {
		return nil, errors.New("error while updating user")
	}
	if err := u.repos.User.SetImage(&updatedUser, payload.Image); err != nil {
		return nil, errors.New("error while updating user image")
	}
	return &payloads.GenericsSuccessFlagResponse{
		Successful: true,
		Message:    "user updated successfully",
	}, nil
}

func (u *User) DeleteUser(ctx context.Context, userID uint64) (*payloads.GenericsSuccessFlagResponse, error) {
	user := domain.GetUserDTO(ctx)
	if user == nil {
		return nil, errors.New("invalid jwt user")
	}
	if user.ID != userID {
		return nil, errors.New("input user id does not match with your id")
	}
	if err := u.repos.User.Delete(&user.User); err != nil {
		return nil, errors.New("error while deleting user")
	}
	return &payloads.GenericsSuccessFlagResponse{
		Successful: true,
		Message:    "user deleted successfully",
	}, nil
}
