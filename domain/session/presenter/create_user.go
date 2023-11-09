package presenter

import (
	"net/http"

	"github.com/aa-ar/budgeter-service/domain/model"
	"github.com/aa-ar/budgeter-service/internal/response"
)

type CreateUserPresenter struct{}

func NewCreateUserPresenter() *CreateUserPresenter {
	return &CreateUserPresenter{}
}

type createUserResponse struct {
	User *model.User `json:"user"`
}

func (p *CreateUserPresenter) PrepareResponse(user *model.User) *response.Response {
	return response.New(http.StatusCreated, createUserResponse{user}, nil)
}
