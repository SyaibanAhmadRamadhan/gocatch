package Jdb

import (
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"

	"github.com/SyaibanAhmadRamadhan/jolly"
)

type DockerTest struct {
	Pool      *dockertest.Pool
	Network   *docker.Network
	Resources []*dockertest.Resource
}

func (d *DockerTest) CleanUp() {
	for _, resource := range d.Resources {
		jolly.PanicIF(d.Pool.Purge(resource))
	}
	jolly.PanicIF(d.Pool.Client.RemoveNetwork(d.Network.ID))
}

func (d *DockerTest) NewContainer(options *dockertest.RunOptions, checkFunc func(res *dockertest.Resource) error) {
	resource, err := d.Pool.RunWithOptions(options)
	jolly.PanicIF(err)

	err = d.Pool.Retry(func() error {
		return checkFunc(resource)
	})
	jolly.PanicIF(err)

	d.Resources = append(d.Resources, resource)
}

func InitDockerTest() *DockerTest {
	pool := NewDockerPool()
	network := NewDockerNetwork(pool)
	return &DockerTest{
		Pool:    pool,
		Network: network,
	}
}

func NewDockerPool() *dockertest.Pool {
	pool, err := dockertest.NewPool("")
	if err != nil {
		panic(err)
	}

	err = pool.Client.Ping()
	if err != nil {
		panic(err)
	}

	return pool
}

func NewDockerNetwork(pool *dockertest.Pool) *docker.Network {
	network, err := pool.Client.CreateNetwork(docker.CreateNetworkOptions{
		Name: "dockertest-" + jolly.RandomString(10),
	})
	if err != nil {
		panic(err)
	}

	return network
}
