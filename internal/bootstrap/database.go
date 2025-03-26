package bootstrap

import (
	"log"

	"github.com/spf13/viper"
	"go.dsig.cn/idev/shortener/internal/dal/db/model"
	"go.dsig.cn/idev/shortener/internal/shared"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

// database 初始化数据库
func database() {
	var err error
	dbType := viper.GetString("database.type")

	switch dbType {
	case "sqlite":
		dataPath := viper.GetString("database.sqlite.path")
		log.Printf("database.sqlite.path: %s\n", dataPath)
		// 使用 modernc.org/sqlite 作为驱动
		dialector := sqlite.Dialector{
			DriverName: "sqlite",
			DSN:        dataPath, // 数据库文件
		}
		// 使用 gorm 内置 SQLite 驱动
		// dialector := sqlite.Open(dataPath)
		shared.GlobalDB, err = gorm.Open(dialector, &gorm.Config{})

	default:
		panic("database type not support: " + dbType)
	}

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	migrate()
}

// migrate 数据库迁移 schema
func migrate() {
	log.Println("migrate")
	shared.GlobalDB.AutoMigrate(&model.Urls{})
}
