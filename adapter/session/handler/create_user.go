package handler

import (
	"encoding/json"
	"net/http"

	"github.com/aa-ar/budgeter-service/adapter/session/errors"
	"github.com/aa-ar/budgeter-service/domain/model"
	"github.com/aa-ar/budgeter-service/internal/response"
	"github.com/aa-ar/budgeter-service/internal/svcerror"
	"github.com/sirupsen/logrus"
)

type CreateUserHandler struct {
	usecase createUserUsecase
}

type createUserUsecase interface {
	CreateUser(*model.Session, string) (*response.Response, error)
}

func NewCreateUserHandler(usecase createUserUsecase) *CreateUserHandler {
	return &CreateUserHandler{
		usecase: usecase,
	}
}

func (h *CreateUserHandler) Path() string {
	return "/session/create_user"
}

func (h *CreateUserHandler) Methods() []string {
	return []string{http.MethodPost}
}

type createUserRequest struct {
	Password string
}

func (h *CreateUserHandler) Handler(w http.ResponseWriter, r *http.Request) error {
	body := json.NewDecoder(r.Body)
	var req *createUserRequest
	err := body.Decode(&req)
	if err != nil {
		logrus.Error(err)
		return svcerror.BadRequestError{}
	}

	sess, err := model.NewSession().FromRequest(r)
	if err != nil {
		return errors.MissingOrInvalidSessionError{}
	}

	if req.Password == "" {
		return errors.MissingPasswordError{}
	}

	// TODO: Validate password strength

	res, err := h.usecase.CreateUser(sess, req.Password)
	if err != nil {
		return err
	}

	res.WriteTo(w)
	return nil
}
