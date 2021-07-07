package global

import (
	"fmt"

	"github.com/goools/tools/servicex"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/driver"
)

type Configure struct {
	MongoDbDriver *driver.MongoDbDriver `json:"mongo_db_driver"`
	Host          string                `json:"host"`
	Port          int                   `json:"port"`
	SwaggerHost   string                `json:"swagger_host"`
}

func (c *Configure) ServerAddress() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

var (
	Config *Configure
)

func init() {
	servicex.SetServiceName("srv-name-list", "..")
	Config = &Configure{}
	servicex.ConfP(Config)

	logrus.Infof("config: %+v", Config)
}
