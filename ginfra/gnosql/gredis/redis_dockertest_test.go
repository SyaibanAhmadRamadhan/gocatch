package gredis

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/redis/rueidis"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra"
)

func TestPostgresDockerTest(t *testing.T) {
	dockerTest := ginfra.InitDockerTest()
	defer dockerTest.CleanUp()

	redisDockerTestConf := RedisDockerTestConf{}

	var client rueidis.Client

	dockerTest.NewContainer(redisDockerTestConf.ImageVersion(dockerTest, ""), func(res *dockertest.Resource) error {
		time.Sleep(2 * time.Second)
		conn, err := redisDockerTestConf.ConnectRueidis(res)
		client = conn
		gcommon.PanicIfError(err)

		return nil
	})

	ctx := context.Background()
	// SET key val NX
	err := client.Do(ctx, client.B().Set().Key("key").Value("val").Nx().Build()).Error()
	gcommon.PanicIfError(err)
	// HGETALL hm
	key, err := client.Do(ctx, client.B().Mget().Key("key").Build()).ToArray()
	gcommon.PanicIfError(err)
	fmt.Println(key)

}
