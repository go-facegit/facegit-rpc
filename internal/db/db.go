package db

import (
	// "database/sql"
	"fmt"
	// "os"
	"strings"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/go-facegit/facegit-rpc/internal/conf"
)

var (
	db  *gorm.DB
	err error
)

var Tables = []interface{}{
	new(User),
}

func Init() error {

	dbUser := conf.Database.User
	dbPwd := conf.Database.Password
	dbHost := conf.Database.Host
	dbName := conf.Database.Name

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true", dbUser, dbPwd, dbHost, dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "facegit_",
			SingularTable: true,
		},
		NowFunc: func() time.Time {
			return time.Now().UTC().Truncate(time.Microsecond)
		},
	})

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// SetMaxIdleConns sets the maximum number of connections in the free connection pool
	sqlDB.SetMaxIdleConns(conf.Database.MaxIdleConns)
	// SetMaxOpenConns sets the maximum number of open database connections.
	sqlDB.SetMaxOpenConns(conf.Database.MaxOpenConns)
	// SetConnMaxLifetime Sets the maximum time that the connection can be reused.
	sqlDB.SetConnMaxLifetime(time.Minute)

	for _, table := range Tables {
		if db.Migrator().HasTable(table) {
			continue
		}

		name := strings.TrimPrefix(fmt.Sprintf("%T", table), "*db.")
		err = db.Migrator().AutoMigrate(table)
		if err != nil {
			return errors.Wrapf(err, "auto migrate %q", name)
		}
	}

	return nil
}

func Ping() error {
	sdb, _ := db.DB()
	return sdb.Ping()
}
