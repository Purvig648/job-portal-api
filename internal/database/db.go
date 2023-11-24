package database

import (
	"fmt"
	"job-application-api/internal/config"
	"job-application-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open(cfg config.DataConfig) (*gorm.DB, error) {
	//database := "host=postgres user=postgres password=admin dbname=portal port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort, cfg.Dbsslmode, cfg.DbTimeZone)
	db, err := gorm.Open(postgres.Open(database), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.Migrator().AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	err = db.Migrator().AutoMigrate(&models.Company{})
	if err != nil {
		return nil, err
	}
	err = db.Migrator().AutoMigrate(&models.Job{}, &models.Location{}, &models.Qualification{}, &models.Shift{}, &models.TechnologyStack{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
