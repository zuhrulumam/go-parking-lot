package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zuhrulumam/doit-test/business/usecase"
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
	// search vehicle
	r.app.Get("/search-vehicle", r.SearchVehicle)

	// available spots
	r.app.Get("/available-spot", r.AvailableSpot)

	r.app.Post("/park", r.Park)

	r.app.Post("/unpark", r.UnPark)
}
