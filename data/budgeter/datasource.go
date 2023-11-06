package budgeter

import (
	"github.com/aa-ar/budgeter-service/internal/data/pg"
	"github.com/jmoiron/sqlx"
)

type BudgeterDataSource struct {
	*sqlx.DB
}

func NewBudgeterDataSource(config *pg.Config) *BudgeterDataSource {
	return &BudgeterDataSource{config.Connect(2, 20)}
}
