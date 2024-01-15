package db

import (
	"log"
	"net"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/go-sql-driver/mysql"
	gorm_mysql "gorm.io/driver/mysql"
)

var db *gorm.DB

type Config struct {
	Host            string `json:"host" yaml:"host"`
	Port            int    `json:"port" yaml:"port"`
	Database        string `json:"database" yaml:"database"`
	Username        string `json:"username" yaml:"username"`
	Password        string `json:"password" yaml:"password"`
	Verbose         bool   `json:"verbose" yaml:"verbose"`
	MaxOpenConns    int    `json:"max_open_conns" yaml:"max_open_conns"`
	MaxIdleConns    int    `json:"max_idle_conns" yaml:"max_idle_conns"`
	ConnMaxLifetime int    `json:"conn_max_lifetime" yaml:"conn_max_lifetime"`
	ConnMaxIdleTime int    `json:"conn_max_idle_time" yaml:"conn_max_idle_time"`
	SSLMode         string `json:"ssl_mode" yaml:"ssl_mode"`
}

func Connect(c Config) (err error) {
	dsnConfig := mysql.NewConfig()
	dsnConfig.User = c.Username
	dsnConfig.Passwd = c.Password
	dsnConfig.Addr = net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
	dsnConfig.DBName = c.Database
	dsnConfig.ParseTime = true
	dsnConfig.InterpolateParams = true
	dsnConfig.Collation = "utf8mb4_unicode_ci"
	dsnConfig.Net = "tcp"
	dsnConfig.Params = map[string]string{
		"charset": "utf8mb4",
	}
	dialect := gorm_mysql.Open(dsnConfig.FormatDSN())

	logrus.Debugf("DSN string : %s", dsnConfig.FormatDSN())

	if db, err = gorm.Open(dialect, &gorm.Config{DisableNestedTransaction: true}); err != nil {
		return errors.Wrap(err, "failed to connect database")
	}

	if c.Verbose {
		db.Logger = logger.New(log.New(logrus.StandardLogger().Out, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      logger.Info,
			Colorful:      true,
		})
	}

	return nil
}

func Conn() *gorm.DB {
	return db
}
