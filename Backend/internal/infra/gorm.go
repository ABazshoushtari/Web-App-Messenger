package infra

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func NewGORMConnection(DSN string) (*gorm.DB, error) {

	pg := postgres.New(postgres.Config{
		DSN:                  DSN,
		PreferSimpleProtocol: true,
	})

	log := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second * 2, // Slow SQL threshold
			LogLevel:                  logger.Error,    // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,            // Disable color
		},
	)

	gormDB, err := gorm.Open(pg, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
			NoLowerCase:   false,
			TablePrefix:   "",
		},
		PrepareStmt:     false,
		Logger:          log,
		CreateBatchSize: 100,
	})

	if err != nil {
		return nil, err
	}
	return gormDB, nil
}
