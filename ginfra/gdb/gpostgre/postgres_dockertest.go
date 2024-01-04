package gpostgre

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest/v3"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gopentelemetry/gotelpgx"
)

type PostgresDockerTestConf struct {
	User     string
	Password string
	Host     string
	Port     int
	DB       string
	SSL      string

	ResourceExpired uint
	pool            *ginfra.DockerTest
	image           string
}

func (p *PostgresDockerTestConf) DBURL() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		p.Host, p.Port, p.User, p.Password, p.DB, p.SSL)
}

func (p *PostgresDockerTestConf) ImageVersion(pool *ginfra.DockerTest, version string) *dockertest.RunOptions {
	p.pool = pool
	p.InitConf(version)

	options := &dockertest.RunOptions{
		Name:       "postgres-" + p.pool.Network.Name,
		Repository: "postgres",
		Tag:        version,
		Env: []string{
			"POSTGRES_PASSWORD=" + p.Password,
			"POSTGRES_USER=" + p.User,
			"POSTGRES_DB=" + p.DB,
			"listen_addresses = '*'",
		},
	}

	return options
}

func (p *PostgresDockerTestConf) ConnectPgx(resource *dockertest.Resource) (conn *pgxpool.Pool, err error) {
	if p.ResourceExpired != 0 {
		resource.Expire(p.ResourceExpired)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")

	port, err := strconv.Atoi(strings.Split(hostAndPort, ":")[1])
	if err != nil {
		return
	}
	p.Host = strings.Split(hostAndPort, ":")[0]
	p.Port = port

	conn = OpenPgxPool(p.DBURL())

	return
}

func (p *PostgresDockerTestConf) ConnectPgxWithOtel(resource *dockertest.Resource) (conn *pgxpool.Pool, err error) {
	if p.ResourceExpired != 0 {
		resource.Expire(p.ResourceExpired)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")

	port, err := strconv.Atoi(strings.Split(hostAndPort, ":")[1])
	if err != nil {
		return
	}
	p.Host = strings.Split(hostAndPort, ":")[0]
	p.Port = port

	config, err := pgxpool.ParseConfig(p.DBURL())
	if err != nil {
		return
	}
	config.ConnConfig.Tracer = gotelpgx.NewTracer()

	conn = OpenPgxPoolWithConfig(config)
	if err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}

	return
}

func (p *PostgresDockerTestConf) ConnectSqlx(resource *dockertest.Resource) (conn *sqlx.DB, err error) {
	if p.ResourceExpired != 0 {
		resource.Expire(p.ResourceExpired)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")

	port, err := strconv.Atoi(strings.Split(hostAndPort, ":")[1])
	if err != nil {
		return
	}
	p.Host = strings.Split(hostAndPort, ":")[0]
	p.Port = port

	conn = OpenConnSqlxPq(p.DBURL())

	return
}

func (p *PostgresDockerTestConf) ConnectSqlxPgx(resource *dockertest.Resource) (conn *sqlx.DB, err error) {
	if p.ResourceExpired != 0 {
		resource.Expire(p.ResourceExpired)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")

	port, err := strconv.Atoi(strings.Split(hostAndPort, ":")[1])
	if err != nil {
		return
	}
	p.Host = strings.Split(hostAndPort, ":")[0]
	p.Port = port

	conn = OpenConnSqlxPgx(p.DBURL())

	return
}

func (p *PostgresDockerTestConf) InitConf(version string) {
	p.User = gcommon.Ternary(p.User != "", p.User, "root")

	p.Password = gcommon.Ternary(p.Password != "", p.Password, "root")

	p.DB = gcommon.Ternary(p.DB != "", p.DB, "dockertest")

	p.image = gcommon.Ternary(version != "", version, "latest")

	p.SSL = gcommon.Ternary(p.SSL != "", p.SSL, "disable")

}
