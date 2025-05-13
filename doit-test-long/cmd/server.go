package cmd

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"github.com/zuhrulumam/doit-test/business/domain"
	"github.com/zuhrulumam/doit-test/business/usecase"
	"github.com/zuhrulumam/doit-test/handler"
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

	app := fiber.New()

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
	handler.Init(handler.Option{
		Uc:  uc,
		App: app,
	})

	log.Println(app.Listen(":8080"))
}

// TODO: Gracefull shutdown
