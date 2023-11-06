package pg

import (
	"fmt"
)

type Config struct {
	Host     string
	Port     string
	Password string
	Username string
	DB       string
	SSL      string
}

func (c *Config) toString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.DB,
		c.SSL,
	)
}
