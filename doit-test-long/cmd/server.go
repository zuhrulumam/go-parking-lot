package cmd

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"github.com/zuhrulumam/doit-test/business/domain"
	"github.com/zuhrulumam/doit-test/business/usecase"
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
)

func run() {

	// init domain
	dom = domain.Init()

	// init usecase
	uc = usecase.Init(dom, usecase.Option{})

	// init rest

	app := fiber.New()

	log.Println(app.Listen(":3000"))
}
