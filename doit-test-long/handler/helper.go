package handler

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zuhrulumam/doit-test/pkg/errors"
	"github.com/zuhrulumam/doit-test/pkg/logger"
)

func (e *rest) compileError(c *fiber.Ctx, err error) error {

	var (
		httpStatus int
		he         string
		code       = errors.ErrCode(err)
		ctx        = c.Locals("ctx").(context.Context)
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

	logger.LogWithCtx(ctx, e.log, err.Error())

	return c.Status(httpStatus).JSON(ErrorResponse{
		HumanError: he,
		DebugError: err.Error(),
		Success:    false,
	})
}
