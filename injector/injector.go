//go:build wireinject
// +build wireinject

package injector

import (
	controllersUser "github.com/backent/fra-golang/controllers/user"
	"github.com/backent/fra-golang/libs"
	repositoriesUser "github.com/backent/fra-golang/repositories/user"
	servicesUser "github.com/backent/fra-golang/services/user"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var UserSet = wire.NewSet(
	controllersUser.NewControllerUserImpl,
	servicesUser.NewServiceUserImpl,
	repositoriesUser.NewRepositoryUserImpl,
)

func InitializeRouter() *httprouter.Router {
	wire.Build(
		libs.NewDatabase,
		libs.NewRouter,
		UserSet,
	)

	return nil
}
