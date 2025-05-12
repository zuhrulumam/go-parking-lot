package handler

import "github.com/zuhrulumam/doit-test/business/usecase"

type Rest interface {
}

type rest struct {
	uc *usecase.Usecase
}

func Init(uc *usecase.Usecase) Rest {
	e := &rest{}

	e.Serve()

	return e
}

func (r rest) Serve() {
	r.uc.Parking.Park()
}
