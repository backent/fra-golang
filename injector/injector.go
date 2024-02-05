//go:build wireinject
// +build wireinject

package injector

import (
	controllersAuth "github.com/backent/fra-golang/controllers/auth"
	controllersDocument "github.com/backent/fra-golang/controllers/document"
	controllersNotification "github.com/backent/fra-golang/controllers/notification"
	controllersRisk "github.com/backent/fra-golang/controllers/risk"
	controllersUser "github.com/backent/fra-golang/controllers/user"
	"github.com/backent/fra-golang/libs"
	"github.com/backent/fra-golang/middlewares"
	repositoriesAuth "github.com/backent/fra-golang/repositories/auth"
	repositoriesDocument "github.com/backent/fra-golang/repositories/document"
	repositoriesNotification "github.com/backent/fra-golang/repositories/notification"
	repositoriesRejectNote "github.com/backent/fra-golang/repositories/rejectnote"
	repositoriesRisk "github.com/backent/fra-golang/repositories/risk"
	repositoriesUser "github.com/backent/fra-golang/repositories/user"
	servicesAuth "github.com/backent/fra-golang/services/auth"
	servicesDocument "github.com/backent/fra-golang/services/document"
	servicesNotification "github.com/backent/fra-golang/services/notification"
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

var DocumentSet = wire.NewSet(
	controllersDocument.NewControllerDocumentImpl,
	servicesDocument.NewServiceDocumentImpl,
	repositoriesDocument.NewRepositoryDocumentImpl,
	middlewares.NewDocumentMiddleware,
)

var NotificationSet = wire.NewSet(
	controllersNotification.NewControllerNotificationImpl,
	servicesNotification.NewServiceNotificationImpl,
	repositoriesNotification.NewRepositoryNotificationImpl,
	middlewares.NewNotificationMiddleware,
)

var AuthSet = wire.NewSet(
	controllersAuth.NewControllerAuthImpl,
	servicesAuth.NewServiceAuthImpl,
	repositoriesAuth.NewRepositoryAuthJWTImpl,
	middlewares.NewAuthMiddleware,
)

var RejectNoteSet = wire.NewSet(
	repositoriesRejectNote.NewRepositoryRejectNote,
)

func InitializeRouter() *httprouter.Router {
	wire.Build(
		libs.NewDatabase,
		libs.NewRouter,
		libs.NewValidator,
		UserSet,
		RiskSet,
		AuthSet,
		DocumentSet,
		RejectNoteSet,
		NotificationSet,
	)

	return nil
}
