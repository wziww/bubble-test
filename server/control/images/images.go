package images

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wziww/bubble-test/common/docker"
	"github.com/wziww/bubble-test/server/lib"
)

// GetALL ...
func GetALL() httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		lib.Send(200, res, req, docker.List(req.Context()))
	}
}

// Search ...
func Search() httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		lib.Send(200, res, req, docker.Search(req.Context()))
	}
}
