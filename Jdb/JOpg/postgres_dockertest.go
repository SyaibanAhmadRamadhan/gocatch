package JOpg

import (
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/dockertest/v3"

	"github.com/SyaibanAhmadRamadhan/jolly/Jdb"
	"github.com/SyaibanAhmadRamadhan/jolly/Jtype"
)

type PostgresDockerTestConf struct {
	PostgresConf    PostgresConf
	ResourceExpired uint
	pool            *Jdb.DockerTest
	image           string
}

func (p *PostgresDockerTestConf) ImageVersion(pool *Jdb.DockerTest, version string) *dockertest.RunOptions {
	p.pool = pool
	p.InitConf(version)

	options := &dockertest.RunOptions{
		Name:       "postgres-" + p.pool.Network.Name,
		Repository: "postgres",
		Tag:        version,
		Env: []string{
			"POSTGRES_PASSWORD=" + p.PostgresConf.Password,
			"POSTGRES_USER=" + p.PostgresConf.User,
			"POSTGRES_DB=" + p.PostgresConf.DB,
			"listen_addresses = '*'",
		},
	}

	return options
}

func (p *PostgresDockerTestConf) Connect(resource *dockertest.Resource) (conn *pgxpool.Pool, err error) {
	if p.ResourceExpired != 0 {
		resource.Expire(p.ResourceExpired)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")

	port, err := strconv.Atoi(strings.Split(hostAndPort, ":")[1])
	if err != nil {
		return
	}
	p.PostgresConf.Host = strings.Split(hostAndPort, ":")[0]
	p.PostgresConf.Port = port

	conn = PgxNewConnection(p.PostgresConf, false)

	return
}

func (p *PostgresDockerTestConf) InitConf(version string) {
	p.PostgresConf.User = Jtype.Ternary(p.PostgresConf.User != "", p.PostgresConf.User, "root")

	p.PostgresConf.Password = Jtype.Ternary(p.PostgresConf.Password != "", p.PostgresConf.Password, "root")

	p.PostgresConf.DB = Jtype.Ternary(p.PostgresConf.DB != "", p.PostgresConf.DB, "dockertest")

	p.image = Jtype.Ternary(version != "", version, "latest")

	p.PostgresConf.SSL = Jtype.Ternary(p.PostgresConf.SSL != "", p.PostgresConf.SSL, "disable")

}
