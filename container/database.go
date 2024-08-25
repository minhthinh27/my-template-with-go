package container

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"my-template-with-go/bootstrap"

	"log"
	"os"
	"time"
)

type IDataBaseProvider interface {
	GetDBMain() *gorm.DB
	GetDBSlave() *gorm.DB
}

type databaseProvider struct {
	dbMain  *gorm.DB
	dbSlave *gorm.DB
}

func NewDatabase(config bootstrap.Database, sugar *zap.SugaredLogger) (IDataBaseProvider, func(), error) {
	var (
		data    = &databaseProvider{}
		cfMain  = config.Main
		cfSlave = config.Slave
	)

	cleanup := func() {
		if data != nil && data.GetDBMain() != nil {
			if sqlDB, err := data.GetDBMain().DB(); err == nil {
				sqlDB.Close()
			}
		}

		if data != nil && data.GetDBSlave() != nil {
			if sqlDB, err := data.GetDBSlave().DB(); err == nil {
				sqlDB.Close()
			}
		}
		sugar.Info("closing the data resources")
	}

	if cfMain.GetHost() != "" {
		mainDB, err := connectMain(cfMain)
		if err != nil {
			return data, cleanup, err
		} else {
			data.dbMain = mainDB
		}
	}

	if cfSlave.GetHost() != "" {
		slaveDB, err := connectSlave(cfSlave)
		if err != nil {
			return data, cleanup, err
		} else {
			data.dbSlave = slaveDB
		}
	}

	return data, cleanup, nil
}

func connectMain(cf bootstrap.Main) (*gorm.DB, error) {
	dsnParams := "?charset=utf8mb4" + "&" +
		"collation=utf8mb4_unicode_ci" + "&" +
		"parseTime=True" + "&" +
		"interpolateParams=True" + "&" +
		"readTimeout=10s" + "&" +
		"timeout=10s" + "&" +
		"writeTimeout=10s"

	dsn := cf.GetUserName() + ":" + cf.GetPassword() +
		"@tcp(" + cf.GetHost() + ":" + cf.GetPort() + ")/" +
		cf.GetName() + dsnParams

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   true,
		DryRun:                                   false,
		PrepareStmt:                              true,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		}),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(cf.GetMaxIdleCon())
	sqlDB.SetMaxOpenConns(cf.GetMaxCon())
	sqlDB.SetConnMaxLifetime(2 * time.Minute)
	sqlDB.SetConnMaxLifetime(3 * time.Minute)

	if sqlDB == nil {
		return nil, errors.New("cannot open connection to database")
	}

	return db, nil
}

func connectSlave(cf bootstrap.Slave) (*gorm.DB, error) {
	dsnParams := "?charset=utf8mb4" + "&" +
		"collation=utf8mb4_unicode_ci" + "&" +
		"parseTime=True" + "&" +
		"interpolateParams=True" + "&" +
		"readTimeout=10s" + "&" +
		"timeout=10s" + "&" +
		"writeTimeout=10s"

	dsn := cf.GetUserName() + ":" + cf.GetPassword() +
		"@tcp(" + cf.GetHost() + ":" + cf.GetPort() + ")/" +
		cf.GetName() + dsnParams

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   true,
		DryRun:                                   false,
		PrepareStmt:                              true,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		}),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(cf.GetMaxIdleCon())
	sqlDB.SetMaxOpenConns(cf.GetMaxCon())
	sqlDB.SetConnMaxLifetime(2 * time.Minute)
	sqlDB.SetConnMaxLifetime(3 * time.Minute)

	if sqlDB == nil {
		return nil, errors.New("cannot open connection to database")
	}

	return db, nil
}

func (p *databaseProvider) GetDBMain() *gorm.DB {
	return p.dbMain.Session(&gorm.Session{SkipHooks: false})
}

func (p *databaseProvider) GetDBSlave() *gorm.DB {
	return p.dbSlave.Session(&gorm.Session{SkipHooks: false})
}
