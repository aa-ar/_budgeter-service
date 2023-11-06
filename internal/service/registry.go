package service

type Registry interface {
	Handlers() []Handler
}
