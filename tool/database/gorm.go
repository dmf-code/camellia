package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
	"sync"
	"time"
)

var once sync.Once
var db *gorm.DB

// Config mysql config.
type Config struct {
	Addr         string        // for trace
	DSN          string        // write data source name.
	ReadDSN      []string      // read data source name.
	Active       int           // pool
	Idle         int           // pool
	IdleTimeout  time.Duration // connect max life time.
	QueryTimeout time.Duration // query sql timeout
	ExecTimeout  time.Duration // execute sql timeout
	TranTimeout  time.Duration // transaction sql timeout
}

func GetInstance() *gorm.DB {
	return db
}

func InitMysql(c *Config)  {
	once.Do(func() {
		var err error
		db, err = connect(c)
		// 初始化 mysql 失败立刻结束进程
		if err != nil {
			panic(err)
		}
	})
}

func connect(c *Config) (*gorm.DB, error) {
	d, err := gorm.Open(mysql.New(mysql.Config{
		DSN: c.DSN, // DSN data source name
		DefaultStringSize: 256, // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "",   // 表名前缀，`User`表为`t_users`
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
			NameReplacer: strings.NewReplacer("CID", "Cid"), // 在转为数据库名称之前，使用NameReplacer更改结构/字段名称。
		},
	})

	if err != nil {
		return nil, err
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := d.DB()

	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(c.Active)
	sqlDB.SetMaxIdleConns(c.Idle)
	sqlDB.SetConnMaxLifetime(time.Duration(c.IdleTimeout))
	return d, nil
}
