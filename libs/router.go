package libs

import (
	controllersUser "github.com/backent/fra-golang/controllers/user"
	"github.com/backent/fra-golang/exceptions"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	controllersUser controllersUser.ControllerUserInterface,
) *httprouter.Router {
	router := httprouter.New()

	router.GET("/users", controllersUser.FindAll)
	router.GET("/users/:id", controllersUser.FindById)
	router.POST("/users", controllersUser.Create)
	router.PUT("/users/:id", controllersUser.Update)
	router.DELETE("/users/:id", controllersUser.Delete)

	router.PanicHandler = exceptions.RouterPanicHandler
	return router
}
