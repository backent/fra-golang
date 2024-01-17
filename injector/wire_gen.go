// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	auth3 "github.com/backent/fra-golang/controllers/auth"
	risk3 "github.com/backent/fra-golang/controllers/risk"
	user3 "github.com/backent/fra-golang/controllers/user"
	"github.com/backent/fra-golang/libs"
	"github.com/backent/fra-golang/middlewares"
	"github.com/backent/fra-golang/repositories/auth"
	"github.com/backent/fra-golang/repositories/risk"
	"github.com/backent/fra-golang/repositories/user"
	auth2 "github.com/backent/fra-golang/services/auth"
	risk2 "github.com/backent/fra-golang/services/risk"
	user2 "github.com/backent/fra-golang/services/user"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

// Injectors from injector.go:

func InitializeRouter() *httprouter.Router {
	db := libs.NewDatabase()
	repositoryUserInterface := user.NewRepositoryUserImpl()
	validate := libs.NewValidator()
	repositoryAuthInterface := auth.NewRepositoryAuthJWTImpl()
	userMiddleware := middlewares.NewUserMiddleware(validate, repositoryUserInterface, repositoryAuthInterface)
	serviceUserInterface := user2.NewServiceUserImpl(db, repositoryUserInterface, userMiddleware)
	controllerUserInterface := user3.NewControllerUserImpl(serviceUserInterface)
	repositoryRiskInterface := risk.NewRepositoryRiskImpl()
	riskMiddleware := middlewares.NewRiskMiddleware(validate, repositoryRiskInterface, repositoryAuthInterface)
	serviceRiskInterface := risk2.NewServiceRiskImpl(db, repositoryRiskInterface, riskMiddleware)
	controllerRiskInterface := risk3.NewControllerRiskImpl(serviceRiskInterface)
	authMiddleware := middlewares.NewAuthMiddleware(validate, repositoryUserInterface, repositoryAuthInterface)
	serviceAuthInterface := auth2.NewServiceAuthImpl(db, repositoryAuthInterface, authMiddleware)
	controllerAuthInterface := auth3.NewControllerAuthImpl(serviceAuthInterface)
	router := libs.NewRouter(controllerUserInterface, controllerRiskInterface, controllerAuthInterface)
	return router
}

// injector.go:

var UserSet = wire.NewSet(user3.NewControllerUserImpl, user2.NewServiceUserImpl, user.NewRepositoryUserImpl, middlewares.NewUserMiddleware)

var RiskSet = wire.NewSet(risk3.NewControllerRiskImpl, risk2.NewServiceRiskImpl, risk.NewRepositoryRiskImpl, middlewares.NewRiskMiddleware)

var AuthSet = wire.NewSet(auth3.NewControllerAuthImpl, auth2.NewServiceAuthImpl, auth.NewRepositoryAuthJWTImpl, middlewares.NewAuthMiddleware)
