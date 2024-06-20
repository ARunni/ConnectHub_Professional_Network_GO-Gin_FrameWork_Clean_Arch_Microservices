package db

import (
	"fmt"

	"github.com/ARunni/ConnetHub_Notification/pkg/config"
	"github.com/ARunni/ConnetHub_Notification/pkg/utils/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBname, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if dbErr != nil {
		return nil, dbErr
	}

	var modelsToMigrate = []interface{}{
		&models.Notification{},
	}

	for _, model := range modelsToMigrate {
		if err := db.AutoMigrate(model); err != nil {
			fmt.Println("Error migrating model:", err)
			return nil, err
		}
	}

	return db, nil
}
