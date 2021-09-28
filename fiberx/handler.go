package fiberx

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gopkg-dev/pkg/errors"
)

var (
	ErrNotFound        = errors.NotFound("NOT_FOUND", "资源不存在")
	ErrTooManyRequests = errors.New(http.StatusTooManyRequests, "TOO_MANY_REQUESTS", "请求过于频繁")
)

func DefaultNotFoundHandler(ctx *fiber.Ctx) error {
	return ErrNotFound
}

func DefaultLimitReachedHandler(ctx *fiber.Ctx) error {
	return ErrTooManyRequests
}

func DefaultErrorHandler(ctx *fiber.Ctx, err error) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	if e, ok := err.(*fiber.Error); ok {
		return ctx.Status(e.Code).JSON(Response{
			Code:        errors.UnknownCode,
			Reason:      errors.UnknownReason,
			Message:     e.Message,
			RequestID:   GetRequestID(ctx),
			RequestTime: time.Now().Unix(),
		})

	}
	e := errors.FromError(err)
	return ctx.Status(e.Code).JSON(Response{
		Code:        e.Code,
		Reason:      e.Reason,
		Message:     e.Message,
		Data:        nil,
		Metadata:    e.Metadata,
		RequestID:   GetRequestID(ctx),
		RequestTime: time.Now().Unix(),
	})
}
