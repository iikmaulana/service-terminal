package controller

import (
	"github.com/iikmaulana/service-terminal/lib"
	"github.com/iikmaulana/service-terminal/models"
	"github.com/iikmaulana/service-terminal/service"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
	"github.com/uzzeet/uzzeet-gateway/libs/utils/uttime"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo service.UserRepo
}

func NewUserUsageUsecase(userRepo service.UserRepo) service.UserUsageUsecase {
	return userUsecase{userRepo: userRepo}
}

func (u userUsecase) UserCreateUsecase(form models.UserCreateRequest) (result string, serr serror.SError) {
	password := []byte(form.Password)
	passwordHash, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	form.Password = string(passwordHash)

	tmpResult, err := u.userRepo.UserCreateRepo(form)
	if err != nil {
		return result, err
	}

	result = tmpResult
	return result, nil
}

func (u userUsecase) UserUpdateUsecase(form models.UserUpdateRequest) (result string, serr serror.SError) {
	tmpResult, err := u.userRepo.UserUpdateRepo(form)
	if err != nil {
		return result, err
	}

	result = tmpResult
	return result, nil
}

func (u userUsecase) UserListUsecase(form models.FilterParams) (result models.UserListResult, serr serror.SError) {
	tmpResult, err := u.userRepo.UserListRepo(form)
	if err != nil {
		return result, err
	}

	result = tmpResult
	return result, nil
}

func (u userUsecase) UserViewUsecase(id string) (result models.UserViewResult, serr serror.SError) {
	tmpResult, err := u.userRepo.UserViewRepo(id)
	if err != nil {
		return result, err
	}

	result = models.UserViewResult{
		Id:        tmpResult.Id,
		Username:  tmpResult.Username,
		CreatedAt: tmpResult.CreatedAt,
		Status:    tmpResult.Status,
		LastLogin: tmpResult.LastLogin,
	}

	return result, nil
}

func (u userUsecase) UserDeleteUsecase(id string) (result string, serr serror.SError) {
	_, err := u.userRepo.UserDeleteRepo(id)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (u userUsecase) UserLoginUsecase(username string, password string) (result string, serr serror.SError) {
	if username == "" || password == "" {
		tmpErr := serror.New("Username atau password tidak tidak terdaftar")
		return result, tmpErr
	}

	tmpResult, err := u.userRepo.UserViewByUseranmeRepo(username)
	if err != nil {
		return result, err
	}

	if lib.CheckPasswordHash(password, tmpResult.Password) {
		tmpResultJwt, errx := lib.GenerateJWT(tmpResult.Id, tmpResult.Username, tmpResult.Status, tmpResult.LastLogin)
		if errx != nil {
			return result, serror.NewFromError(errx)
		}
		result = tmpResultJwt

		_, err := u.userRepo.UserUpdateRepo(models.UserUpdateRequest{
			Id:        tmpResult.Id,
			LastLogin: uttime.Now().Format("2006-01-02 15:04:05"),
		})

		if err != nil {
			return result, serror.NewFromError(errx)
		}
	} else {
		tmpErr := serror.New("Username atau password tidak tidak sesuai")
		return result, tmpErr
	}

	return result, nil
}
