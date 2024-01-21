package libs

import (
	controllersAuth "github.com/backent/fra-golang/controllers/auth"
	controllersDocument "github.com/backent/fra-golang/controllers/document"
	controllersRisk "github.com/backent/fra-golang/controllers/risk"
	controllersUser "github.com/backent/fra-golang/controllers/user"
	"github.com/backent/fra-golang/exceptions"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	controllersUser controllersUser.ControllerUserInterface,
	controllersRisk controllersRisk.ControllerRiskInterface,
	controllersAuth controllersAuth.ControllerAuthInterface,
	controllersDocument controllersDocument.ControllerDocumentInterface,
) *httprouter.Router {
	router := httprouter.New()

	router.POST("/login", controllersAuth.Login)
	router.POST("/register", controllersAuth.Register)

	router.GET("/users", controllersUser.FindAll)
	router.GET("/users/:id", controllersUser.FindById)
	router.POST("/users", controllersUser.Create)
	router.PUT("/users/:id", controllersUser.Update)
	router.DELETE("/users/:id", controllersUser.Delete)

	router.GET("/risks", controllersRisk.FindAll)
	router.GET("/risks/:id", controllersRisk.FindById)
	router.POST("/risks", controllersRisk.Create)
	router.PUT("/risks/:id", controllersRisk.Update)
	router.DELETE("/risks/:id", controllersRisk.Delete)

	router.GET("/documents", controllersDocument.FindAll)
	router.GET("/documents/:id", controllersDocument.FindById)
	router.POST("/documents", controllersDocument.Create)
	router.PUT("/documents/:id", controllersDocument.Update)
	router.DELETE("/documents/:id", controllersDocument.Delete)

	router.GET("/documents-product-name", controllersDocument.GetProductDistinct)

	router.PanicHandler = exceptions.RouterPanicHandler
	return router
}
