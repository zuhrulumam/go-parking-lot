package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zuhrulumam/doit-test/pkg/errors"
)

func (e *rest) compileError(c *fiber.Ctx, err error) error {

	code := errors.ErrCode(err)
	var httpStatus int

	switch code {
	case 400:
		httpStatus = http.StatusBadRequest
	default:
		httpStatus = http.StatusInternalServerError
	}

	return c.Status(httpStatus).JSON(fiber.Map{
		"sys":         err.Error(),
		"human_error": err,
	})
}
