package verifyInterfaceCompliance

import "net/http"

type HandlerG struct {
	// ...
}

var _ http.Handler = (*HandlerG)(nil)

func (h *HandlerG) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	// ...
}

type LogHandler struct {
	h http.Handler
	//log *zap.Logger
}

var _ http.Handler = LogHandler{}

func (h LogHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	// ...
}
