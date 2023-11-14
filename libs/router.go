package libs

import (
	controllersDocument "github.com/backent/fra-golang/controllers/document"
	controllersUser "github.com/backent/fra-golang/controllers/user"
	"github.com/backent/fra-golang/exceptions"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	controllersUser controllersUser.ControllerUserInterface,
	controllersDocument controllersDocument.ControllerDocumentInterface,
) *httprouter.Router {
	router := httprouter.New()

	router.GET("/users", controllersUser.FindAll)
	router.GET("/users/:id", controllersUser.FindById)
	router.POST("/users", controllersUser.Create)
	router.PUT("/users/:id", controllersUser.Update)
	router.DELETE("/users/:id", controllersUser.Delete)

	router.GET("/documents", controllersDocument.FindAll)
	router.GET("/documents/:id", controllersDocument.FindById)
	router.POST("/documents", controllersDocument.Create)
	router.PUT("/documents/:id", controllersDocument.Update)
	router.DELETE("/documents/:id", controllersDocument.Delete)

	router.PanicHandler = exceptions.RouterPanicHandler
	return router
}
