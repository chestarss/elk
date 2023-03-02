// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/mailru/easyjson"
	"github.com/chestarss/elk/internal/uuid/ent"
)

// Basic HTTP Error Response
type ErrResponse struct {
	Code   int         `json:"code"`             // http response status code
	Status string      `json:"status"`           // user-level status message
	Errors interface{} `json:"errors,omitempty"` // application-level error
}

func (e ErrResponse) MarshalToHTTPResponseWriter(w http.ResponseWriter) (int, error) {
	d, err := easyjson.Marshal(e)
	if err != nil {
		return 0, err
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(d)))
	w.WriteHeader(e.Code)
	return w.Write(d)
}

func BadRequest(w http.ResponseWriter, msg interface{}) (int, error) {
	return ErrResponse{
		Code:   http.StatusBadRequest,
		Status: http.StatusText(http.StatusBadRequest),
		Errors: msg,
	}.MarshalToHTTPResponseWriter(w)
}

func Conflict(w http.ResponseWriter, msg interface{}) (int, error) {
	return ErrResponse{
		Code:   http.StatusConflict,
		Status: http.StatusText(http.StatusConflict),
		Errors: msg,
	}.MarshalToHTTPResponseWriter(w)
}

func Forbidden(w http.ResponseWriter, msg interface{}) (int, error) {
	return ErrResponse{
		Code:   http.StatusForbidden,
		Status: http.StatusText(http.StatusForbidden),
		Errors: msg,
	}.MarshalToHTTPResponseWriter(w)
}

func InternalServerError(w http.ResponseWriter, msg interface{}) (int, error) {
	return ErrResponse{
		Code:   http.StatusInternalServerError,
		Status: http.StatusText(http.StatusInternalServerError),
		Errors: msg,
	}.MarshalToHTTPResponseWriter(w)
}

func NotFound(w http.ResponseWriter, msg interface{}) (int, error) {
	return ErrResponse{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
		Errors: msg,
	}.MarshalToHTTPResponseWriter(w)
}

func Unauthorized(w http.ResponseWriter, msg interface{}) (int, error) {
	return ErrResponse{
		Code:   http.StatusUnauthorized,
		Status: http.StatusText(http.StatusUnauthorized),
		Errors: msg,
	}.MarshalToHTTPResponseWriter(w)
}

type (
	// User3451555716View represents the data serialized for the following serialization group combinations:
	// []
	User3451555716View struct {
		ID   int       `json:"id,omitempty"`
		UUID uuid.UUID `json:"uuid,omitempty"`
	}
	User3451555716Views []*User3451555716View
)

func NewUser3451555716View(e *ent.User) *User3451555716View {
	if e == nil {
		return nil
	}
	return &User3451555716View{
		ID:   e.ID,
		UUID: e.UUID,
	}
}

func NewUser3451555716Views(es []*ent.User) User3451555716Views {
	if len(es) == 0 {
		return nil
	}
	r := make(User3451555716Views, len(es))
	for i, e := range es {
		r[i] = NewUser3451555716View(e)
	}
	return r
}
