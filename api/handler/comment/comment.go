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
		// Get all comments.
		c, err := a.CommentService.GetCommentsAll()
		if err != nil {
			response.InternalServerError(w, err)
			return
		}

		// Successful response.
		response.APIResponse(w, c)
		return
	})
}

func CreateComment(a *app.App) request.HandlerFunc {
	return request.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the vars.
		vars := mux.Vars(r)
		userID, err := request.GetVarInt64(vars, "user_id")
		if err != nil {
			response.InternalServerError(w, err)
			return
		}

		comment, err := request.GetVarString(vars, "comment")
		if err != nil {
			response.InternalServerError(w, err)
			return
		}

		// Create the comment.
		c, err := a.CommentService.CreateComment(userID, comment)
		if err != nil {
			response.InternalServerError(w, err)
			return
		}

		// Successful response.
		response.APIResponse(w, c)
		return
	})
}

func DeleteComment(a *app.App) request.HandlerFunc {
	return request.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get the vars.
		vars := mux.Vars(r)
		userID, err := request.GetVarInt64(vars, "user_id")
		if err != nil {
			response.InternalServerError(w, err)
			return
		}

		commentID, err := request.GetVarInt64(vars, "comment_id")
		if err != nil {
			response.InternalServerError(w, err)
			return
		}

		// Delete the comment.
		err = a.CommentService.DeleteComment(userID, commentID)
		if err != nil {
			response.InternalServerError(w, err)
			return
		}

		// Successful response.
		response.APIResponse(w, true)
		return

	})
}

func LikeComment(a *app.App) request.HandlerFunc {
	return request.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get the vars.
		vars := mux.Vars(r)
		userID, err := request.GetVarInt64(vars, "user_id")
		if err != nil {
			response.InternalServerError(w, err)
			return
		}

		commentID, err := request.GetVarInt64(vars, "comment_id")
		if err != nil {
			response.InternalServerError(w, err)
			return
		}

		// Delete the comment.
		err = a.CommentService.CreateCommentLike(userID, commentID)
		if err != nil {
			response.InternalServerError(w, err)
			return
		}

		// Successful response.
		response.APIResponse(w, true)
		return
	})
}
