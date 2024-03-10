//go:build wireinject
// +build wireinject

package injector

import (
	controllersAuth "github.com/backent/fra-golang/controllers/auth"
	controllersDashboard "github.com/backent/fra-golang/controllers/dashboard"
	controllersDocument "github.com/backent/fra-golang/controllers/document"
	controllersNotification "github.com/backent/fra-golang/controllers/notification"
	controllersRisk "github.com/backent/fra-golang/controllers/risk"
	controllersUser "github.com/backent/fra-golang/controllers/user"
	controllersUserRegistration "github.com/backent/fra-golang/controllers/user_registration"
	"github.com/backent/fra-golang/libs"
	"github.com/backent/fra-golang/middlewares"
	repositoriesAuth "github.com/backent/fra-golang/repositories/auth"
	repositoriesDocument "github.com/backent/fra-golang/repositories/document"
	repositoriesNotification "github.com/backent/fra-golang/repositories/notification"
	repositoriesRejectNote "github.com/backent/fra-golang/repositories/rejectnote"
	repositoriesRisk "github.com/backent/fra-golang/repositories/risk"
	repositoriesUser "github.com/backent/fra-golang/repositories/user"
	repositoriesUserRegistration "github.com/backent/fra-golang/repositories/user_registration"
	servicesAuth "github.com/backent/fra-golang/services/auth"
	servicesDashboard "github.com/backent/fra-golang/services/dashboard"
	servicesDocument "github.com/backent/fra-golang/services/document"
	servicesNotification "github.com/backent/fra-golang/services/notification"
	servicesRisk "github.com/backent/fra-golang/services/risk"
	servicesUser "github.com/backent/fra-golang/services/user"
	servicesUserRegistration "github.com/backent/fra-golang/services/user_registration"
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
	repositoriesDocument.NewRepositoryDocumentSearchEsImpl,
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

var UserRegistrationSet = wire.NewSet(
	controllersUserRegistration.NewControllerUserRegistrationImpl,
	servicesUserRegistration.NewServiceUserRegistrationImpl,
	repositoriesUserRegistration.NewRepositoryUserRegistrationImpl,
	middlewares.NewUserRegistrationMiddleware,
)

var DashboardSet = wire.NewSet(
	controllersDashboard.NewControllerDashboardImpl,
	servicesDashboard.NewServiceDashboardImpl,
	middlewares.NewDashboardMiddleware,
)

func InitializeRouter() *httprouter.Router {
	wire.Build(
		libs.NewDatabase,
		libs.NewRouter,
		libs.NewValidator,
		libs.NewElastic,
		UserSet,
		RiskSet,
		AuthSet,
		DocumentSet,
		RejectNoteSet,
		NotificationSet,
		UserRegistrationSet,
		DashboardSet,
	)

	return nil
}
