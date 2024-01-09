package web

import (
	"net/http"
	"strconv"
)

type RequestPagination interface {
	SetSkip(skip int)
	SetTake(take int)
	GetSkip() int
	GetTake() int
}

func SetPagination(request RequestPagination, r *http.Request) {
	if r.URL.Query().Has("take") {
		take, err := strconv.Atoi(r.URL.Query().Get("take"))
		if err != nil {
			panic(err)
		}
		request.SetTake(take)
	} else {
		request.SetTake(10)
	}

	if r.URL.Query().Has("skip") {
		skip, err := strconv.Atoi(r.URL.Query().Get("skip"))
		if err != nil {
			panic(err)
		}
		request.SetSkip(skip)
	} else {
		request.SetSkip(0)
	}

}
