package services

/*
   Services are the Use cases of the clean Architecture
*/

import (
	"user/constants"
	"user/models"
	"user/utils/jwt"

	"golang.org/x/crypto/bcrypt"
)

//UserRepository :
type UserRepository interface {
	FindUserByID(userID int64) (models.User, error)
	FindUserByEmail(email string) (models.User, error)
	CreateUser(models.SignupBodyParams) (models.User, error)
	FindUserByIDAndUpdate(userID int64, params models.UpdateProfileBodyParams) error
}

//UserService :
type UserService interface {
	SignupUser(models.SignupBodyParams) (models.AuthResponse, error)
	SigninUser(models.SigninBodyParams) (models.AuthResponse, error)

	GetProfile(userID int64) (models.ProfileResponse, error)
	UpdateProfile(userID int64, params models.UpdateProfileBodyParams) (models.ProfileResponse, error)
}

type userService struct {
	repo UserRepository
}

//NewUserService :
func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) SignupUser(params models.SignupBodyParams) (models.AuthResponse, error) {

	response := models.AuthResponse{}

	hashPWD, err := bcrypt.GenerateFromPassword([]byte(params.Password), 8)
	if err != nil {
		return response, err
	}

	params.Password = string(hashPWD)
	user, err := s.repo.CreateUser(params)
	//For simplicity i am just returning email exist error
	//but error should be return based on the error code
	if err != nil {
		return response, constants.ErrorEmailExist
	}

	response.Token = jwt.CreateTokenWithClaims(user.ID)
	return response, err
}

func (s *userService) SigninUser(params models.SigninBodyParams) (models.AuthResponse, error) {

	response := models.AuthResponse{}
	user, err := s.repo.FindUserByEmail(params.Email)
	//For simplicity i am just returning email not exist error
	//but error should be return based on the error code
	if err != nil {
		return response, constants.ErrorEmailNotExist
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		return response, constants.ErrorInvalidPassword
	}

	response.Token = jwt.CreateTokenWithClaims(user.ID)
	return response, err
}

func (s *userService) GetProfile(userID int64) (models.ProfileResponse, error) {

	var response models.ProfileResponse

	user, err := s.repo.FindUserByID(userID)
	//error should be handled properly based on error code.
	if err != nil {
		return response, err
	}

	//send only required params to user
	response.Email = user.Email
	response.Name = user.Name

	return response, nil
}

func (s *userService) UpdateProfile(userID int64,
	params models.UpdateProfileBodyParams) (models.ProfileResponse, error) {

	var response models.ProfileResponse

	//if password present in the params and then replace it
	// with the hashed password
	if len(params.Password) > 0 {
		hashPWD, err := bcrypt.GenerateFromPassword([]byte(params.Password), 8)
		if err != nil {
			return response, err
		}
		params.Password = string(hashPWD)
	}

	err := s.repo.FindUserByIDAndUpdate(userID, params)
	if err != nil {
		return response, err
	}
	return s.GetProfile(userID)
}
