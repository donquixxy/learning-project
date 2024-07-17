package cmd

import (
	"context"
	"fmt"
	"learning-project/config"
	"learning-project/internal/app"
	"learning-project/internal/module/user"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	driversql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	cobra.OnInitialize()

	migrateCmd.Flags().Int("version", 0, "Migration version to run (0 to run all up)")
	migrateCmd.Flags().String("file", "", "File name to migrate")
	LearningCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		version, _ := cmd.Flags().GetInt("version")
		file, _ := cmd.Flags().GetString("file")
		if file != "" {
			return runSpecificMigration(file)
		}
		return runMigrations(version)
	},
}

var LearningCmd = &cobra.Command{
	Use:   "app",
	Short: "App",
	Run: func(cmd *cobra.Command, args []string) {
		appConfig := config.GetAppConfiguration()

		log.Printf("Starting Application: %v", appConfig.Name)
		log.Printf("At Environment: %v", appConfig.AppEnv)

		router := app.NewRouter()
		runApp(router)

		go func() {
			if err := router.Echo.Start(fmt.Sprintf(":%v", appConfig.AppPort)); err != nil {
				log.Fatalf("error starting app: %v", err)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		signal.Notify(quit, syscall.SIGTERM)
		<-quit
		router.Echo.Shutdown(context.Background())

	},
}

func runMigrations(version int) error {

	dbCfg := config.GetDatabaseConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?parseTime=true", dbCfg.Username, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.LearningName)

	db, err := gorm.Open(driversql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("error getting *sql.DB from GORM: %v", err)
	}

	driver, err := mysql.WithInstance(sqlDB, &mysql.Config{})
	if err != nil {
		log.Fatalf("error creating migrate driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("error creating migrate instance: %v", err)
	}

	if version > 0 {
		err = m.Steps(version)
	} else {
		err = m.Up()
	}

	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error applying migrations: %w", err)
	}

	log.Println("Migrations applied successfully!")
	return nil
}

func runSpecificMigration(file string) error {
	dbCfg := config.GetDatabaseConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?parseTime=true", dbCfg.Username, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.LearningName)

	db, err := gorm.Open(driversql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	// Read SQL commands from file
	sqlCommands, err := os.ReadFile(fmt.Sprintf("migrations/%s", file))
	if err != nil {
		return fmt.Errorf("error reading migration file: %w", err)
	}

	// Execute SQL commands
	err = db.Exec(string(sqlCommands)).Error
	if err != nil {
		return fmt.Errorf("error executing migration file: %w", err)
	}

	log.Printf("%v Migration executed successfully", file)
	return nil
}

func runApp(router *app.Router) {
	user.InitUserModule(router)
}
