package database

import (
	"errors"
	"fmt"
	"github.com/MicroMolekula/auth-service/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	dbConn, err := gorm.Open(postgres.Open(makeConfigString(cfg)), &gorm.Config{})
	if err != nil {
		return nil, errors.New(fmt.Sprint("database connection error: ", err))
	}

	return dbConn, nil
}

func makeConfigString(cfg *config.Config) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DbName,
		cfg.Database.Timezone,
	)
}
