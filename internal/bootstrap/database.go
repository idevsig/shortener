package bootstrap

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"go.dsig.cn/shortener/internal/dal/db/model"
	"go.dsig.cn/shortener/internal/shared"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

// database 初始化数据库
func database() {
	var err error
	var dialector gorm.Dialector

	dbType := os.Getenv("DATABASE_TYPE")
	if dbType == "" {
		dbType = viper.GetString("database.type")
	}

	switch dbType {
	case "sqlite":
		dsn := viper.GetString("database.sqlite.path")
		log.Printf("database.sqlite: %s\n", dsn)
		// 使用 modernc.org/sqlite 作为驱动
		dialector = sqlite.Dialector{
			DriverName: "sqlite",
			DSN:        dsn, // 数据库文件
		}
		// 使用 gorm 内置 SQLite 驱动
		// dialector := sqlite.Open(dsn)

	case "mysql":
		host := viper.GetString("database.mysql.host")
		if host == "" {
			host = "localhost"
		}
		port := viper.GetInt("database.mysql.port")
		if port == 0 {
			port = 3306
		}
		charset := viper.GetString("database.mysql.charset")
		if charset == "" {
			charset = "utf8mb4"
		}
		parseTime := viper.GetBool("database.mysql.parse_time")
		loc := viper.GetString("database.mysql.loc")
		if loc == "" {
			loc = "Local"
		}
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
			viper.GetString("database.mysql.user"),
			viper.GetString("database.mysql.password"),
			host,
			port,
			viper.GetString("database.mysql.database"),
			charset,
			parseTime,
			loc,
		)
		log.Printf("database.mysql: %s\n", dsn)
		// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
		dialector = mysql.Open(dsn)

	case "postgres":
		host := viper.GetString("database.postgres.host")
		if host == "" {
			host = "localhost"
		}
		port := viper.GetInt("database.postgres.port")
		if port == 0 {
			port = 5432
		}
		sslmode := viper.GetString("database.postgres.sslmode")
		if sslmode == "" {
			sslmode = "disable"
		}
		loc := viper.GetString("database.postgres.timezone")
		if loc == "" {
			loc = "Asia/Shanghai"
		}
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
			host,
			viper.GetString("database.postgres.user"),
			viper.GetString("database.postgres.password"),
			viper.GetString("database.postgres.database"),
			port,
			sslmode,
			loc,
		)
		log.Printf("database.postgres: %s\n", dsn)
		dialector = postgres.Open(dsn)

	default:
		panic("database type not support: " + dbType)
	}

	shared.GlobalDB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	migrate()
}

// migrate 数据库迁移 schema
func migrate() {
	log.Println("migrate")
	err := shared.GlobalDB.AutoMigrate(&model.Urls{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}
}
