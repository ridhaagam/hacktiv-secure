package adapter

import (
	userService "secure/challenge-2/application/usecases/user"
	userRepository "secure/challenge-2/infrastructure/repository/postgres/user"
	userController "secure/challenge-2/infrastructure/restapi/controllers/user"

	"gorm.io/gorm"
)

// UserAdapter is a function that returns a user controller
func UserAdapter(db *gorm.DB) *userController.Controller {
	uRepository := userRepository.Repository{DB: db}
	service := userService.Service{UserRepository: uRepository}
	return &userController.Controller{UserService: service}
}
