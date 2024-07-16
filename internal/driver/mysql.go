package driver

import (
	"fmt"
	"learning-project/config"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	gorm_sql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dbConfig := config.GetDatabaseConfig()

	location, err := time.LoadLocation("Asia/Makassar")

	if err != nil {
		log.Fatalf("Could not load database location: %v", err)
	}

	addr := fmt.Sprintf("%v:%v", dbConfig.Host, dbConfig.Port)

	cfgs := mysql.Config{
		User:                 dbConfig.Username,
		Passwd:               dbConfig.Password,
		Net:                  "tcp",
		Addr:                 addr,
		DBName:               dbConfig.LearningName,
		AllowNativePasswords: true,
		ParseTime:            true,
		Loc:                  location,
	}

	cfg := gorm_sql.Config{
		DSNConfig: &cfgs,
	}

	gormDb, err := gorm.Open(gorm_sql.New(cfg), &gorm.Config{})

	if err != nil {
		log.Fatalf("[InitDatabase] - Failed to init database !")
	}

	db, err := gormDb.DB()

	if err != nil {
		log.Fatalf("[InitDatabase] - Failed to initialize database")
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("[InitDatabase] - Database  failed to connect")
	}

	log.Println("[InitDatabase] - Connected to database")
	return gormDb
}
