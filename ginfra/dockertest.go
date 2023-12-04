package ginfra

import (
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
)

type DockerTest struct {
	Pool      *dockertest.Pool
	Network   *docker.Network
	Resources []*dockertest.Resource
}

func (d *DockerTest) CleanUp() {
	for _, resource := range d.Resources {
		gcommon.PanicIfError(d.Pool.Purge(resource))
	}
	gcommon.PanicIfError(d.Pool.Client.RemoveNetwork(d.Network.ID))
}

func (d *DockerTest) NewContainer(options *dockertest.RunOptions, checkFunc func(res *dockertest.Resource) error) {
	resource, err := d.Pool.RunWithOptions(options)
	gcommon.PanicIfError(err)

	err = d.Pool.Retry(func() error {
		return checkFunc(resource)
	})
	gcommon.PanicIfError(err)

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
		Name: "dockertest-" + gcommon.RandomAlphabeticString(10),
	})
	if err != nil {
		panic(err)
	}

	return network
}
