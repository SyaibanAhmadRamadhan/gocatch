package gmongodb

import (
	"fmt"
	"strings"

	"github.com/ory/dockertest/v3"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra"
)

type MongoDockerTestConf struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string

	ResourceExpired uint
	pool            *ginfra.DockerTest
	image           string
}

func (m *MongoDockerTestConf) URI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/?replicaSet=rs0", m.Username, m.Password, m.Host, m.Port)
}

func (m *MongoDockerTestConf) ImageVersion(pool *ginfra.DockerTest, version string) *dockertest.RunOptions {
	m.pool = pool
	m.InitConf(version)
	return &dockertest.RunOptions{
		Repository: `mongo`,
		Name:       `dockertest-mongo-` + m.pool.Network.Name,
		Tag:        m.image,
		NetworkID:  pool.Network.ID,
		Env: []string{
			`MONGO_INITDB_ROOT_USERNAME=` + m.Username,
			`MONGO_INITDB_ROOT_PASSWORD=` + m.Password,
		},
		Cmd: []string{"--replSet", "rs0", "--bind_ip_all"},
	}
}

func (m *MongoDockerTestConf) ConnectClient(resource *dockertest.Resource) (conn *mongo.Client, err error) {
	if m.ResourceExpired != 0 {
		resource.Expire(m.ResourceExpired)
	}

	hostAndPort := resource.GetHostPort("27017/tcp")
	port := strings.Split(hostAndPort, ":")[1]
	m.Host = strings.Split(hostAndPort, ":")[0]
	m.Port = port

	conn, err = OpenConnMongoClient(m.URI())
	gcommon.PanicIfError(err)

	return
}

func (m *MongoDockerTestConf) ConnectDB(resource *dockertest.Resource) (db *mongo.Database, err error) {
	if m.ResourceExpired != 0 {
		resource.Expire(m.ResourceExpired)
	}

	hostAndPort := resource.GetHostPort("27017/tcp")
	port := strings.Split(hostAndPort, ":")[1]
	m.Host = strings.Split(hostAndPort, ":")[0]
	m.Port = port

	db, err = OpenConnMongoDB(m.URI(), m.Database)
	gcommon.PanicIfError(err)

	return
}

func (m *MongoDockerTestConf) InitConf(version string) {
	m.Username = gcommon.Ternary(m.Username != "", m.Username, "root")

	m.Password = gcommon.Ternary(m.Password != "", m.Password, "root")

	m.Database = gcommon.Ternary(m.Database != "", m.Database, "dockertest")

	m.image = gcommon.Ternary(version != "", version, "latest")
}
