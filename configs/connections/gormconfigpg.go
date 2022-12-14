package connections

import (
	"fmt"
	"os"
	"sesi7/cmd/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionGormPg() (*gorm.DB, error) {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file1")
	// }

	pqsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("HOST_PQSQL"), os.Getenv("PORT_PQSQL"), os.Getenv("USER_PQSQL"), os.Getenv("PASS_PQSQL"), os.Getenv("DBNAME_PQSQL"))

	db, err := gorm.Open(postgres.Open(pqsqlInfo), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.Debug().AutoMigrate(entities.Employee{})

	return db, nil
}
