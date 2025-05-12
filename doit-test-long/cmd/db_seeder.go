package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ParkingSpot struct {
	ID       uint `gorm:"primaryKey"`
	Floor    int
	Row      int
	Col      int
	Type     string `gorm:"size:1"` // 'B', 'M', 'A', 'X'
	Active   bool
	Occupied bool

	Vehicle Vehicle `gorm:"foreignKey:SpotID"`
}

type Vehicle struct {
	ID            uint   `gorm:"primaryKey"`
	VehicleNumber string `gorm:"uniqueIndex"`
	VehicleType   string `gorm:"size:1"` // 'B', 'M', 'A'
	SpotID        uint   // Foreign key to ParkingSpot
	ParkedAt      time.Time
	UnparkedAt    *time.Time
}

var seedCommand = &cobra.Command{
	Use:  "seed [floors] [rows] [cols]",
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		var floors, rows, cols int
		_, _ = fmt.Sscanf(args[0], "%d", &floors)
		_, _ = fmt.Sscanf(args[1], "%d", &rows)
		_, _ = fmt.Sscanf(args[2], "%d", &cols)

		db, _ := connectDB()

		seed(db, floors, rows, cols)
	},
}

func seed(db *gorm.DB, floors, rows, cols int) {
	// migrate db
	if err := db.AutoMigrate(&ParkingSpot{}, &Vehicle{}); err != nil {
		log.Fatalf("failed to migrate tables: %v", err)
	}

	var spots []ParkingSpot

	for f := 1; f <= floors; f++ {
		for r := 1; r <= rows; r++ {
			for c := 1; c <= cols; c++ {
				t := randomType()
				spot := ParkingSpot{
					Floor:  f,
					Row:    r,
					Col:    c,
					Type:   t,
					Active: t != "X",
				}
				spots = append(spots, spot)
			}
		}
	}

	db.CreateInBatches(spots, 1000)
}

func connectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func randomType() string {
	types := []string{"B", "M", "A", "X"}
	return types[rand.Intn(len(types))]
}
