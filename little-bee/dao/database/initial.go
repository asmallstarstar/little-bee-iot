package database

import (
	"dao/serialization"
	"time"
	myerror "util/error"
	"util/i18n"
	"util/log"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB = nil

func InitDB(dsn string, logLevel int) {
	var err error

	zapGormLogger := New(
		&ZapWriter{},
		Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.LogLevel(logLevel),
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: zapGormLogger.LogMode(logger.Info),
	})

	if err != nil {
		msg := i18n.GetLocalTextByDefaultLocal(myerror.MessageCodeValueEnum_desc_id[int32(myerror.ERROR_OPEN_DATABAE)], nil)
		log.Logger.Fatal(msg, zap.String(myerror.ERROR_DESC_STRING, err.Error()))
	}

	schema.RegisterSerializer("pbtimestamp", serialization.PBTimestampSerializer{})
}
