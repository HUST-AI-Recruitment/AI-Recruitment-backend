package model

import (
	"AI-Recruitment-backend/internal/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBEngine(databaseConfig *config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseConfig.UserName,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.DBName,
		databaseConfig.Charset,
		databaseConfig.ParseTime,
	)

	//if global.Config.App.Debug == true {
	//	config.Logger = logger.Default.LogMode(logger.Info)
	//}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(databaseConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseConfig.MaxOpenConns)

	return db, nil
}

func MigrateSchema(db *gorm.DB, schemas []interface{}) error {
	err := db.AutoMigrate(schemas...)
	if err != nil {
		return err
	}
	return nil
}
