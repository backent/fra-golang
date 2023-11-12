package user

import (
	"net/http"

	servicesUser "github.com/backent/fra-golang/services/user"
	"github.com/julienschmidt/httprouter"
)

type ControllerUserImpl struct {
	servicesUser.ServiceUserInterface
}

func NewControllerUserImpl(servicesUser servicesUser.ServiceUserInterface) ControllerUserInterface {
	return &ControllerUserImpl{
		ServiceUserInterface: servicesUser,
	}
}

func (implementation *ControllerUserImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
func (implementation *ControllerUserImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
func (implementation *ControllerUserImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
func (implementation *ControllerUserImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
func (implementation *ControllerUserImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
