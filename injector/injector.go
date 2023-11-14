//go:build wireinject
// +build wireinject

package injector

import (
	controllersDocument "github.com/backent/fra-golang/controllers/document"
	controllersUser "github.com/backent/fra-golang/controllers/user"
	"github.com/backent/fra-golang/libs"
	"github.com/backent/fra-golang/middlewares"
	repositoriesDocument "github.com/backent/fra-golang/repositories/document"
	repositoriesUser "github.com/backent/fra-golang/repositories/user"
	servicesDocument "github.com/backent/fra-golang/services/document"
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

var DocumentSet = wire.NewSet(
	controllersDocument.NewControllerDocumentImpl,
	servicesDocument.NewServiceDocumentImpl,
	repositoriesDocument.NewRepositoryDocumentImpl,
	middlewares.NewDocumentMiddleware,
)

func InitializeRouter() *httprouter.Router {
	wire.Build(
		libs.NewDatabase,
		libs.NewRouter,
		libs.NewValidator,
		UserSet,
		DocumentSet,
	)

	return nil
}
