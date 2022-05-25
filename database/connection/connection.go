package connection

import (
	"gorm.io/gorm"
)

type Connection interface {
	connect() *gorm.DB
}

func Connect(connection Connection) *gorm.DB {
	return connection.connect()
}
