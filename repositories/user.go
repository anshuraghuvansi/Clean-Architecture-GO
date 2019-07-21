package repositories

/*
   Clean Architecture : Repositories are the outermost layer is generally composed of frameworks and tools such as the Database
*/

import (
	"user/models"
	"user/services"
	"user/utils/database"
)

type userRepository struct {
	db *database.DataBase
}

//NewUserRepository :
func NewUserRepository(db *database.DataBase) services.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(params models.SignupBodyParams) (models.User, error) {
	user := params.ToEntity()
	result := r.db.Create(&user)
	return user, result.Error
}

func (r *userRepository) FindUserByEmail(email string) (models.User, error) {

	var user models.User
	result := r.db.Where(models.User{Email: email}).Take(&user)
	return user, result.Error
}

func (r *userRepository) FindUserByID(userID int64) (models.User, error) {

	var user models.User
	result := r.db.Where(models.User{ID: userID}).Take(&user)
	return user, result.Error
}

func (r *userRepository) FindUserByIDAndUpdate(userID int64, params models.UpdateProfileBodyParams) error {
	return r.db.Model(models.User{}).Where(models.User{ID: userID}).Updates(params).Error
}
