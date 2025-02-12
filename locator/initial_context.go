package locator

import (
	"goproject/database"
	"goproject/repository"
	"goproject/service"
)

type InitialContext struct{}

func (initialContext *InitialContext) Lookup(serviceName string) service.Service {
	switch serviceName {
	case "UserService":
		userRepository := repository.NewUserRepository(database.DB)
		return service.NewUserService(userRepository)
	default:
		return nil
	}
}
