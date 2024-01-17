//go:build wireinject
// +build wireinject

package injector

import (
	controllersAuth "github.com/backent/fra-golang/controllers/auth"
	controllersRisk "github.com/backent/fra-golang/controllers/risk"
	controllersUser "github.com/backent/fra-golang/controllers/user"
	"github.com/backent/fra-golang/libs"
	"github.com/backent/fra-golang/middlewares"
	repositoriesAuth "github.com/backent/fra-golang/repositories/auth"
	repositoriesRisk "github.com/backent/fra-golang/repositories/risk"
	repositoriesUser "github.com/backent/fra-golang/repositories/user"
	servicesAuth "github.com/backent/fra-golang/services/auth"
	servicesRisk "github.com/backent/fra-golang/services/risk"
	servicesUser "github.com/backent/fra-golang/services/user"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var UserSet = wire.NewSet(
	controllersUser.NewControllerUserImpl,
	servicesUser.NewServiceUserImpl,
	repositoriesUser.NewRepositoryUserImpl,
	middlewares.NewUserMiddleware,
)

var RiskSet = wire.NewSet(
	controllersRisk.NewControllerRiskImpl,
	servicesRisk.NewServiceRiskImpl,
	repositoriesRisk.NewRepositoryRiskImpl,
	middlewares.NewRiskMiddleware,
)

var AuthSet = wire.NewSet(
	controllersAuth.NewControllerAuthImpl,
	servicesAuth.NewServiceAuthImpl,
	repositoriesAuth.NewRepositoryAuthJWTImpl,
	middlewares.NewAuthMiddleware,
)

func InitializeRouter() *httprouter.Router {
	wire.Build(
		libs.NewDatabase,
		libs.NewRouter,
		libs.NewValidator,
		UserSet,
		RiskSet,
		AuthSet,
	)

	return nil
}
