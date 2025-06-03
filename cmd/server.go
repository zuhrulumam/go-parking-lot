package cmd

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"github.com/zuhrulumam/go-parking-lot/business/domain"
	"github.com/zuhrulumam/go-parking-lot/business/usecase"
	"github.com/zuhrulumam/go-parking-lot/handler"
	"github.com/zuhrulumam/go-parking-lot/pkg/logger"
	"github.com/zuhrulumam/go-parking-lot/pkg/middlewares"
	"go.uber.org/zap"
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
	lg  *zap.Logger
)

func run() {

	lg = logger.NewZapLogger()

	app := fiber.New()
	app.Use(middlewares.RequestContextMiddleware(lg))

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
		Log: lg,
	})

	log.Println(app.Listen(":8080"))
}

// TODO: Gracefull shutdown
