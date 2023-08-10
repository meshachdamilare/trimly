package postgres

import (
	"context"
	"fmt"
	"github/meshachdamilare/trimly/model"
	"github/meshachdamilare/trimly/settings/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type Postgres struct {
	db *gorm.DB
}

var (
	db           *gorm.DB
	queryTimeout = 30 * time.Second
)

func ConnectToDB() *gorm.DB {
	fmt.Println(dsn())
	database, err := gorm.Open(postgres.Open(dsn()), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not connect to postgres, got error: %s", err)
	}
	db = database
	if err := migrateDB(); err != nil {
		log.Fatalf("could not run db migrations, got error: %s", err)
	}
	// IF EVERYTHING IS OKAY, THEN CONNECTION IS ESTABLISHED
	fmt.Println("POSTGRES CONNECTION ESTABLISHED")
	return db
}

func GetDB() *Postgres {
	return &Postgres{
		db: db,
	}
}

func (p *Postgres) DBWithTimeout(contx context.Context) (*gorm.DB, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(contx, queryTimeout)
	return p.db.WithContext(ctx), cancel
}

// dsn returns string for the postgres db connection
func dsn() string {
	pgHost := config.Config.DBHost
	pgPort := config.Config.DBPort
	pgUser := config.Config.DBUser
	pgDBName := config.Config.DBName
	pgPassword := config.Config.DBPassword

	dsn := "host=" + pgHost + " user=" + pgUser +
		" password=" + pgPassword + " dbname=" + pgDBName + " port=" + pgPort + " sslmode=disable"

	return dsn
}

// migrateDB creates db schemas
func migrateDB() error {
	err := db.AutoMigrate(
		&model.URL{},
		//	&model.User{},
	)
	if err != nil {
		return err
	}
	return nil
}
