package connection

import (
	. "Duna/database/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type MySqlConnection struct {
	Config Config
}

func (sql MySqlConnection) connect() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		sql.Config.User,
		sql.Config.Password,
		sql.Config.Host,
		sql.Config.Port,
		sql.Config.DBName,
	)
	db, _err := gorm.Open(mysql.Open(dsn))

	if _err != nil {
		log.Fatalf(`Unable to connect to database`)
	}

	return db
}
