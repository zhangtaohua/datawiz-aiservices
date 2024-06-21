package bootstrap

import (
	"errors"
	"fmt"
	"time"

	"datawiz-aiservices/pkg/config"
	"datawiz-aiservices/pkg/database"
	"datawiz-aiservices/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "gorm.io/gorm/logger"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {

	var dbConfig gorm.Dialector
	databaseType := config.Get("database.connection")
	switch databaseType {
	case "postgres":
		check_or_creat_pg_database(config.Get("database.postgres.database"))
		// 构建 DSN 信息
		// dsn format,  <username>:<password>@<hostname>:<port>/<db>?[k=v& ......]
		// dsn: "postgres://postgres:123456@localhost:5432/public?sslmode=disable&TimeZone=Asia/Shanghai"
		dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable&TimeZone=Asia/Shanghai",
			config.Get("database.postgres.username"),
			config.Get("database.postgres.password"),
			config.Get("database.postgres.host"),
			config.Get("database.postgres.port"),
			config.Get("database.postgres.database"),
		)
		fmt.Printf("dsn = %v", dsn)
		dbConfig = postgres.Open(dsn)
	case "mysql":
		// 构建 DSN 信息
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.Get("database.mysql.username"),
			config.Get("database.mysql.password"),
			config.Get("database.mysql.host"),
			config.Get("database.mysql.port"),
			config.Get("database.mysql.database"),
			config.Get("database.mysql.charset"),
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	case "sqlite":
		// 初始化 sqlite
		database := config.Get("database.sqlite.database")
		dbConfig = sqlite.Open(database)
	default:
		panic(errors.New("database connection not supported"))
	}

	// 连接数据库，并设置 GORM 的日志模式
	// database.Connect(dbConfig, logger.Default.LogMode(logger.Info))
	database.Connect(dbConfig, logger.NewGormLogger())

	// 解决 解决错误: 函数 uuid_generate_v4() 不存在
	if databaseType == "postgres" {
		database.SetPostgresExtension("pgcrypto")
		database.SetPostgresExtension("uuid-ossp")
	}

	// 设置最大连接数
	database.SQLDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	database.SQLDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)
}

func check_or_creat_pg_database(databasename string) bool {
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable&TimeZone=Asia/Shanghai",
		config.Get("database.postgres.username"),
		config.Get("database.postgres.password"),
		config.Get("database.postgres.host"),
		config.Get("database.postgres.port"),
		"postgres",
	)
	fmt.Printf("dsn postgres = %v \n", dsn)
	dbConfig := postgres.Open(dsn)
	database.Connect(dbConfig, logger.NewGormLogger())

	var exists bool
	raw_str := fmt.Sprintf("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = '%v')", databasename)
	fmt.Printf("dsn raw_str = %v \n", raw_str)

	row := database.DB.Raw(raw_str).Row()
	fmt.Printf("dsn row data = %v \n", row)
	if err := row.Scan(&exists); err != nil {
		logger.Fatal("DATABASE Error ", zap.Any("error", err))
	}
	fmt.Printf("dsn exists = %v \n", exists)

	if exists {
		fmt.Println("Database aidb already exists.")
	} else {
		// 创建 aidb 数据库
		if err := database.DB.Exec("CREATE DATABASE aidb").Error; err != nil {
			logger.Fatal("DATABASE Error ", zap.Any("error", err))
			return false
		}
		fmt.Println("Database aidb created successfully.")
	}
	return true
}
