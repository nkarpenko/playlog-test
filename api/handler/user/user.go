package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nkarpenko/playlog-test/api/app"
	"github.com/nkarpenko/playlog-test/api/common/request"
	"github.com/nkarpenko/playlog-test/api/common/response"
)

// Auth errors
var (
	errNotImplemented = response.APIError{Code: 0, Message: "This method has not been implemented yet", Status: http.StatusNotImplemented}
)

// Handlers all of request
func Handlers(r *mux.Router, a *app.App) {
	// @router business
	s := r.PathPrefix("/user").Subrouter()

	s.HandleFunc("/login/{name}", Login(a)).Methods("GET", "OPTIONS")
}

func Login(a *app.App) request.HandlerFunc {
	return request.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=ascii")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		// Get the vars.
		vars := mux.Vars(r)
		name, err := request.GetVarString(vars, "name")
		if err != nil {
			response.InternalServerError(w, err)
			return
		}

		// Login the user.
		u, err := a.UserService.EnsureUserExists(name)
		if err != nil {
			response.InternalServerError(w, err)
			return
		}

		// Successful response.
		response.APIResponse(w, u)
		return
	})
}
