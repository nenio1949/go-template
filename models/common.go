package models

import (
	"database/sql/driver"
	"fmt"
	"go-template/global"
	"os"
	"strconv"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

// 自增ID主键
type Model struct {
	ID        int       `json:"id" gorm:"primaryKey;comment:主键"`
	CreatedAt LocalTime `json:"created_at" gorm:"autoCreateTime:milli;comment:创建时间"`
	UpdatedAt LocalTime `json:"updated_at" gorm:"autoUpdateTime:milli;comment:更新时间"`
	Deleted   int       `json:"deleted" gorm:"index;default:0;comment:是否逻辑删除"`
}

/*****************时间格式化***************/
type LocalTime struct {
	time.Time
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(t.Time)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t.Time)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

/*******************时间格式化****************/

// 初始化数据库
func InitializeDB() *gorm.DB {
	// 根据驱动配置进行初始化
	switch global.App.Config.Database.Driver {
	case "mysql":
		return initMySqlGorm()
	default:
		return initMySqlGorm()
	}
}

// 初始化 mysql gorm.DB
func initMySqlGorm() *gorm.DB {
	dbConfig := global.App.Config.Database

	if dbConfig.Database == "" {
		return nil
	}
	dsn := dbConfig.UserName + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" +
		dbConfig.Database + "?charset=" + dbConfig.Charset + "&parseTime=True&loc=Local&clientFoundRows=true"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if database, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // 表前缀
			SingularTable: true, // 禁用表名复数
		},
		DisableForeignKeyConstraintWhenMigrating: false,                               // 不禁用自动创建外键约束
		Logger:                                   logger.Default.LogMode(logger.Info), // 使用自定义 Logger
	}); err != nil {
		global.App.Log.Error("mysql connect failed, err:", zap.Any("err", err))
		return nil
	} else {
		db = database
		sqlDB, _ := database.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
		initMySqlTables(database)
		return database
	}
}

// func getGormLogger() logger.Interface {
// 	var logMode logger.LogLevel

// 	switch global.App.Config.Database.LogMode {
// 	case "silent":
// 		logMode = logger.Silent
// 	case "error":
// 		logMode = logger.Error
// 	case "warn":
// 		logMode = logger.Warn
// 	case "info":
// 		logMode = logger.Info
// 	default:
// 		logMode = logger.Info
// 	}

// 	return logger.New(getGormLogWriter(), logger.Config{
// 		SlowThreshold:             200 * time.Millisecond,                          // 慢 SQL 阈值
// 		LogLevel:                  logMode,                                         // 日志级别
// 		IgnoreRecordNotFoundError: false,                                           // 忽略ErrRecordNotFound（记录未找到）错误
// 		Colorful:                  !global.App.Config.Database.EnableFileLogWriter, // 禁用彩色打印
// 	})
// }

// // 自定义 gorm Writer
// func getGormLogWriter() logger.Writer {
// 	var writer io.Writer

// 	// 是否启用日志文件
// 	if global.App.Config.Database.EnableFileLogWriter {
// 		// 自定义 Writer
// 		writer = &lumberjack.Logger{
// 			Filename:   global.App.Config.Log.RootDir + "/" + global.App.Config.Database.LogFilename,
// 			MaxSize:    global.App.Config.Log.MaxSize,
// 			MaxBackups: global.App.Config.Log.MaxBackups,
// 			MaxAge:     global.App.Config.Log.MaxAge,
// 			Compress:   global.App.Config.Log.Compress,
// 		}
// 	} else {
// 		// 默认 Writer
// 		writer = os.Stdout
// 	}
// 	return log.New(writer, "\r\n", log.LstdFlags)
// }

// 数据库表初始化
func initMySqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		User{},
		Role{},
		Department{},
	)
	if err != nil {
		global.App.Log.Error("migrate table failed", zap.Any("err", err))
		os.Exit(0)
	}
}
