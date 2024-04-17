package dao

import (
	"context"
	"gin-mall/conf/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"strings"
	"time"
)

var _db *gorm.DB

func InitMysql() {
	mConfig := sql.Config.MySql["default"]
	pathRead := strings.Join([]string{mConfig.UserName, ":", mConfig.Password, "@tcp(", mConfig.DbHost, ":", mConfig.DbPort, ")/", mConfig.DbName, "?charset=" + mConfig.Charset + "&parseTime=true"}, "")
	pathWrite := strings.Join([]string{mConfig.UserName, ":", mConfig.Password, "@tcp(", mConfig.DbHost, ":", mConfig.DbPort, ")/", mConfig.DbName, "?charset=" + mConfig.Charset + "&parseTime=true"}, "")

	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       pathRead,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameColumn:   true,
		DontSupportRenameIndex:    true,
		SkipInitializeWithVersion: false, // 根据版本自动控制
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	// 获取一个连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(20)  // 设置连接池，空闲
	sqlDB.SetMaxIdleConns(100) // 打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db

	_ = _db.Use(dbresolver.Register(
		dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(pathRead)},
			Replicas: []gorm.Dialector{mysql.Open(pathWrite), mysql.Open(pathWrite)},
			Policy:   dbresolver.RandomPolicy{},
		}))

	_db.Set("gorm:table_options", "charset=utf8mb4")
	err = migrate()
	if err != nil {
		panic(err)
	}
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
