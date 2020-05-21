package lib

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/wziww/bubble-test/common/logger"
)

var (
	// XBubbleCode http_status_code
	XBubbleCode = "status"
)

// Send ...
func Send(code int, w http.ResponseWriter, req *http.Request, v interface{}) {
	if w.Header().Get("X-Bubble") != "true;" {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Bubble", "true;")
		w.Header().Set(XBubbleCode, strconv.Itoa(code))
		w.WriteHeader(code)
		response, _ := json.Marshal(v)
		_, err := w.Write(response)
		if err != nil {
			if err.Error() == "http: connection has been hijacked" {
				return
			}
			logger.Error(err.Error())
		}
	}
}
