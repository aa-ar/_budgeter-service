package svcerror

type Error interface {
	Status() int
	Error() string
}
