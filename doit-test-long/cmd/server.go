package cmd

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"github.com/zuhrulumam/doit-test/business/domain"
	"github.com/zuhrulumam/doit-test/business/usecase"
	"gorm.io/gorm"
)

var serverCommand = &cobra.Command{
	Use:   "start",
	Short: "start server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

var (
	dom *domain.Domain
	uc  *usecase.Usecase
	db  *gorm.DB
)

func run() {

	// init sql
	g, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}

	db = g

	// init domain
	dom = domain.Init(domain.Option{
		DB: db,
	})

	// init usecase
	uc = usecase.Init(dom, usecase.Option{})

	// init rest

	app := fiber.New()

	log.Println(app.Listen(":3000"))
}

// TODO: Gracefull shutdown
