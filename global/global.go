package global

import (
	"github.com/goools/tools/servicex"
	"github.com/sirupsen/logrus"
	"github.com/siskinc/srv-name-list/driver"
)

type Configure struct {
	MongoDbDriver *driver.MongoDbDriver `json:"mongo_db_driver"`
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
