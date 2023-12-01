package JOLpg

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/SyaibanAhmadRamadhan/jolly"
)

type PostgresConf struct {
	User     string
	Password string
	Host     string
	Port     int
	DB       string
	SSL      string
}

func (p PostgresConf) DBURL() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		p.Host, p.Port, p.User, p.Password, p.DB, p.SSL)
}

type Adapter struct {
	*pgxpool.Pool
}

func PgxNewConnection(pgConf PostgresConf, withPing bool) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	conn, err := pgxpool.New(ctx, pgConf.DBURL())
	jolly.PanicIF(err)

	if withPing {
		err = conn.Ping(ctx)
		jolly.PanicIF(err)
	}

	return conn
}
