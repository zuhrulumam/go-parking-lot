package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/zuhrulumam/doit-test/business/usecase"
	_ "github.com/zuhrulumam/doit-test/docs" // replace with your module
)

type Rest interface {
}

type Option struct {
	Uc  *usecase.Usecase
	App *fiber.App
}

type rest struct {
	uc  *usecase.Usecase
	app *fiber.App
}

func Init(opt Option) Rest {
	e := &rest{
		uc:  opt.Uc,
		app: opt.App,
	}

	e.Serve()

	return e
}

func (r rest) Serve() {
	// swagger
	r.app.Get("/swagger/*", swagger.HandlerDefault)

	// search vehicle
	r.app.Get("/vehicle/search", r.SearchVehicle)

	// available spots
	r.app.Get("/spot/available", r.AvailableSpot)

	r.app.Post("/vehicle/park", r.Park)

	r.app.Post("/vehicle/unpark", r.UnPark)
}
