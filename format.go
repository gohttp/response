package response

import "net/http"
import "bitbucket.org/ww/goautoneg"

func Format(r *http.Request, handlers map[string]func()) {
	contentTypes := []string{}
	for k := range handlers {
		contentTypes = append(contentTypes, k)
	}

	accept := r.Header.Get("Accept")

	contentType := goautoneg.Negotiate(accept, contentTypes)

	if handler, ok := handlers[contentType]; ok {
		handler()
	} else if handler, ok := handlers["default"]; ok {
		handler()
	}
}
