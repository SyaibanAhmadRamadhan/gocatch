package JOLpg

import (
	"context"
	"testing"

	"github.com/ory/dockertest/v3"

	"github.com/SyaibanAhmadRamadhan/jolly/Jdb"
)

func TestPostgresDockerTest(t *testing.T) {
	dockerTest := Jdb.InitDockerTest()
	defer dockerTest.CleanUp()

	postgresDockerTest := PostgresDockerTestConf{}

	dockerTest.NewContainer(postgresDockerTest.ImageVersion(dockerTest, ""), func(res *dockertest.Resource) error {
		conn, err := postgresDockerTest.Connect(res)

		err = conn.Ping(context.Background())

		return err
	})
}
