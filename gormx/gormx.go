package gormx

import (
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Config 配置参数
type Config struct {
	Debug                                    bool   // 是否开启调试模式
	DBType                                   string // 数据库类型,mysql,postgres,sqlite
	DSN                                      string // 数据库链接字符串
	TablePrefix                              string // 表名前缀
	MaxLifetime                              int    // 连接最长存活期,超过这个时间连接将不再被复用
	MaxIdleTime                              int    // 设置连接空闲的最大时间
	MaxOpenConns                             int    // 数据库最大连接数
	MaxIdleConns                             int    // 最大空闲连接数
	DisableForeignKeyConstraintWhenMigrating bool   // 迁移时禁用外键约束
}

// NewDB 创建DB实例
func NewDB(cfg *Config) (*gorm.DB, error) {

	var dialector gorm.Dialector

	switch strings.ToLower(cfg.DBType) {
	case "mysql":
		dialector = mysql.Open(cfg.DSN)
	case "postgres":
		dialector = postgres.Open(cfg.DSN)
	default:
		dialector = sqlite.Open(cfg.DSN)
	}

	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   cfg.TablePrefix,
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: cfg.DisableForeignKeyConstraintWhenMigrating,
	}

	db, err := gorm.Open(dialector, config)
	if err != nil {
		return nil, err
	}

	if cfg.Debug {
		db = db.Debug()
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDb.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDb.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Second)
	sqlDb.SetConnMaxIdleTime(time.Duration(cfg.MaxIdleTime) * time.Second)

	return db, nil
}

// AutoMigrate 自动映射数据表
func AutoMigrate(db *gorm.DB, dst ...interface{}) error {
	if db.Dialector.Name() == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	}
	return db.AutoMigrate(dst...)
}
