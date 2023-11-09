package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Response struct {
	Status  int
	Body    interface{}
	Cookies []*http.Cookie
}

type errorBody struct {
	Error   string      `json:"error"`
	Details interface{} `json:"details"`
}

func New(status int, body interface{}, cookies []*http.Cookie) *Response {
	if err, ok := body.(error); ok {
		if err != nil {
			body = errorBody{
				Error:   err.Error(),
				Details: err,
			}
		}
	}
	return &Response{
		Status:  status,
		Body:    body,
		Cookies: cookies,
	}
}

func (r *Response) Marshal() []byte {
	j, err := json.Marshal(r.Body)
	if err != nil {
		logrus.Error(err)
	}
	return j
}

func (r *Response) SetCookies(w http.ResponseWriter) {
	for _, cookie := range r.Cookies {
		http.SetCookie(w, cookie)
	}
}

func (r *Response) WriteTo(w http.ResponseWriter) {
	j := r.Marshal()
	r.SetCookies(w)
	w.WriteHeader(int(r.Status))
	fmt.Fprint(w, string(j))
}
