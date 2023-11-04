package main

import (
	"github.com/aa-ar/budgeter-service/internal/service"
)

func main() {
	service.NewService(3030).Start()
}
