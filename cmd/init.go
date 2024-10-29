package main

import (
	"AI-Recruitment-backend/internal/config"
	"AI-Recruitment-backend/internal/global"
	"AI-Recruitment-backend/internal/model"
)

func initConfig() error {
	// Load the configuration file
	var err error
	global.Config, err = config.NewConfig()
	if err != nil {
		return err
	}
	return nil
}

func initDB() error {
	// Initialize the database
	var err error
	global.DBEngine, err = model.NewDBEngine(&global.Config.Database)
	if err != nil {
		return err
	}

	// Migrate the schema
	err = model.MigrateSchema(global.DBEngine, []interface{}{
		&model.User{},
		&model.Job{},
	})
	if err != nil {
		return err
	}
	return nil
}
