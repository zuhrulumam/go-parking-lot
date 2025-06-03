package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"gorm.io/gorm"
)

var cleanerCommand = &cobra.Command{
	Use: "clean-db",
	Run: func(cmd *cobra.Command, args []string) {

		db, _ := connectDB()

		clean(db)
	},
}

func clean(db *gorm.DB) {

	// migrate db
	if err := db.Migrator().DropTable(&ParkingSpot{}, &Vehicle{}); err != nil {
		log.Fatalf("failed to migrate tables: %v", err)
	}
}
