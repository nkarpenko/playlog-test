package comment

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
	s := r.PathPrefix("/comment").Subrouter()

	s.HandleFunc("/all", GetAllComments(a)).Methods("GET")

	s.HandleFunc("/create", CreateComment(a)).Methods("POST")
	s.HandleFunc("/delete", DeleteComment(a)).Methods("POST")
	s.HandleFunc("/like", LikeComment(a)).Methods("POST")
}

func GetAllComments(a *app.App) request.HandlerFunc {
	return request.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func CreateComment(a *app.App) request.HandlerFunc {
	return request.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func DeleteComment(a *app.App) request.HandlerFunc {
	return request.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func LikeComment(a *app.App) request.HandlerFunc {
	return request.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

// GetDomainAvailability handler
// func GetDomainAvailability(a *app.App) request.HandlerFunc {
// 	return request.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Get the vars.
// 		vars := mux.Vars(r)
// 		name, err := request.GetVarString(vars, common.VarName)
// 		if err != nil {
// 			response.InternalServerError(w, err)
// 			return
// 		}

// 		// Get the domain availability.
// 		p, err := a.DomainService.GetDomainAvailability(name)
// 		if err != nil {
// 			// Domain level errors.
// 			switch err {
// 			case derr.ErrDomainInvalid:
// 				response.APIResponseError(w, apierr.ErrDomainInvalid, err)
// 				return
// 			case derr.ErrDomainAvCheck:
// 				response.APIResponseError(w, apierr.ErrDomainAvCheck, err)
// 				return
// 			case derr.ErrDomainWhoisReq:
// 				response.APIResponseError(w, apierr.ErrDomainWhoisReq, err)
// 				return
// 			default:
// 				response.InternalServerError(w, err)
// 				return
// 			}
// 		}

// 		// Successful response.
// 		response.APIResponse(w, p)
// 		return
// 	})
// }
