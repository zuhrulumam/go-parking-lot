package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zuhrulumam/doit-test/pkg/errors"
)

func (e *rest) compileError(c *fiber.Ctx, err error) error {

	var (
		httpStatus int
		he         string
		code       = errors.ErrCode(err)
	)
	switch code {
	case 400:
		httpStatus = http.StatusBadRequest
		he = errors.EM.Message("EN", "badrequest")
	case 404:
		httpStatus = http.StatusNotFound
		he = errors.EM.Message("EN", "notfound")
	default:
		httpStatus = http.StatusInternalServerError
		he = errors.EM.Message("EN", "internal")
	}

	return c.Status(httpStatus).JSON(fiber.Map{
		"human_error": he,
		"debug_error": err.Error(),
	})
}
