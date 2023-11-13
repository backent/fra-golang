package exceptions

import (
	"net/http"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/web"
	"github.com/go-playground/validator/v10"
)

func RouterPanicHandler(w http.ResponseWriter, r *http.Request, i interface{}) {
	var response web.WebResponse

	if err, ok := i.(validator.ValidationErrors); ok {
		response = web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
	} else if err, ok := i.(BadRequestError); ok {
		response = web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
	} else if err, ok := i.(NotFoundError); ok {
		response = web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   err.Error(),
		}
	} else if err, ok := i.(error); ok {
		response = web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL STATUS SERVER ERROR",
			Data:   err.Error(),
		}
	} else {
		response = web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL STATUS SERVER ERROR",
			Data:   i,
		}
	}

	helpers.ReturnReponseJSON(w, response)
}
