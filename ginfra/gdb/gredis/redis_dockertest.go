package gredis

import (
	"strconv"
	"strings"

	"github.com/ory/dockertest/v3"
	"github.com/redis/rueidis"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra"
)

type RedisDockerTestConf struct {
	Host     string
	Port     int
	DB       int
	Password string

	ResourceExpired uint
	pool            *ginfra.DockerTest
	image           string
}

func (p *RedisDockerTestConf) ImageVersion(pool *ginfra.DockerTest, version string) *dockertest.RunOptions {
	p.pool = pool
	p.InitConf(version)

	options := &dockertest.RunOptions{
		Name:       "redis-" + p.pool.Network.Name,
		Repository: "redis",
		Tag:        version,
		Env:        []string{},
		Cmd: []string{
			"redis-server",
			"--requirepass",
			p.Password,
		},
	}

	return options
}

func (p *RedisDockerTestConf) ConnectRueidis(resource *dockertest.Resource) (adapter *RueidisAdapter, err error) {
	if p.ResourceExpired != 0 {
		resource.Expire(p.ResourceExpired)
	}

	hostAndPort := resource.GetHostPort("6379/tcp")

	port, err := strconv.Atoi(strings.Split(hostAndPort, ":")[1])
	if err != nil {
		return
	}
	p.Host = strings.Split(hostAndPort, ":")[0]
	p.Port = port

	adapter = OpenConnRueidis(rueidis.ClientOption{
		InitAddress: []string{hostAndPort},
		Password:    p.Password,
		SelectDB:    p.DB,
	})

	return
}

func (p *RedisDockerTestConf) InitConf(version string) {
	p.Password = gcommon.Ternary(p.Password != "", p.Password, "root")

	p.DB = gcommon.Ternary(p.DB != 0, p.DB, 0)

	p.image = gcommon.Ternary(version != "", version, "latest")
}
