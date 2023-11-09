package registry

import (
	"github.com/aa-ar/budgeter-service/adapter/session/handler"
	"github.com/aa-ar/budgeter-service/data/budgeter"
	"github.com/aa-ar/budgeter-service/data/session"
	"github.com/aa-ar/budgeter-service/domain/session/presenter"
	"github.com/aa-ar/budgeter-service/domain/session/usecase"
	"github.com/aa-ar/budgeter-service/internal/service"
)

type Registry struct {
	handlers []service.Handler
}

func NewSessionRegistry(s session.SessionDataSource, b budgeter.BudgeterDataSource) *Registry {
	return &Registry{
		handlers: []service.Handler{
			handler.NewSessionInitHandler(usecase.NewSessionInitUsecase(s, b, presenter.NewSessionInitPresenter())),
			handler.NewCreateUserHandler(usecase.NewCreateUserUsecase(s, b, presenter.NewCreateUserPresenter())),
		},
	}
}

func (registry *Registry) Handlers() []service.Handler {
	return registry.handlers
}
