package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func (c *Config) checkConnectionDetails() {
	if c.Host == "" {
		logrus.Fatal("No host provided for PG connection")
	}
	if c.Port == "" {
		logrus.Fatal("No port provided for PG connection")
	}
	if c.Password == "" {
		logrus.Fatal("No password provided for PG connection")
	}
	if c.Username == "" {
		logrus.Fatal("No username provided for PG connection")
	}
	if c.DB == "" {
		logrus.Fatal("No database name provided for PG connection")
	}
	if c.SSL == "" {
		c.SSL = "disable"
	}
}

func (c *Config) Connect(maxIdle int, maxOpen int) *sqlx.DB {
	c.checkConnectionDetails()
	db, err := sqlx.Open("postgres", c.toString())
	if err != nil {
		logrus.Fatal(err)
	}
	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(maxOpen)
	return db
}
