package httpext

import (
	"net/http"
	"net/http/httputil"
)

func Dump(req *http.Request, res *http.Response, verbose bool) {
	if req != nil {
		b, _ := httputil.DumpRequest(req, verbose)
		println(string(b))
	}
	if res != nil {
		b, _ := httputil.DumpResponse(res, verbose)
		println(string(b))
	}
}
