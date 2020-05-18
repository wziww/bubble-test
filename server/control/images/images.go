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
		lib.Send(200, res, req, docker.ImagesGet(req.Context()))
	}
}

// Search ...
func Search() httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		lib.Send(200, res, req, docker.ImagesSearch(req.Context()))
	}
}

// Pull ...
func Pull() httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		lib.Send(200, res, req, docker.ImagesPull(req.Context(), req.URL.Query().Get("image")))
	}
}
