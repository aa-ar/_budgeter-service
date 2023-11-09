package main

import (
	"os"

	_ "github.com/lib/pq"

	budgeterDataSource "github.com/aa-ar/budgeter-service/data/budgeter"
	sessionDataSource "github.com/aa-ar/budgeter-service/data/session"
	session "github.com/aa-ar/budgeter-service/domain/session/registry"
	"github.com/aa-ar/budgeter-service/internal/data/pg"
	"github.com/aa-ar/budgeter-service/internal/data/redis"
	"github.com/aa-ar/budgeter-service/internal/service"
)

var registries []service.Registry

func init() {
	sessionDataSource := sessionDataSource.NewSessionDataSource(
		&redis.Config{
			Addr:     os.Getenv("RDB_ADDR"),
			Password: os.Getenv("RDB_PWD"),
		},
	)

	budgeterDataSource := budgeterDataSource.NewBudgeterDataSource(&pg.Config{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		Password: os.Getenv("PG_PWD"),
		Username: os.Getenv("PG_USER"),
		DB:       os.Getenv("PG_DB"),
		SSL:      os.Getenv("PG_SSL"),
	})

	registries = []service.Registry{
		session.NewSessionRegistry(*sessionDataSource, *budgeterDataSource),
	}
}

func main() {
	service.NewService(3030).
		AttachRegistries(registries).
		Start()
}
