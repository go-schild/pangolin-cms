package cms

import "net/http"

func buildAdminHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("admin panel is currently under development"))
	})
}
