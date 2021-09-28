package fiberx

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gopkg-dev/pkg/gormx"
)

const (
	RequestIDKey = "request_id"
	UserIDKey    = "user_id"
)

// Map is a shortcut for map[string]interface{}, useful for JSON returns
type Map map[string]interface{}

// GetRequestID ...
func GetRequestID(ctx *fiber.Ctx) string {
	return ctx.Locals(RequestIDKey).(string)
}

// SetUserID ...
func SetUserID(ctx *fiber.Ctx, userID int) {
	ctx.Locals(UserIDKey, userID)
}

// GetUserID ...
func GetUserID(ctx *fiber.Ctx) int {
	return ctx.Locals(UserIDKey).(int)
}

// ParseBody is helper function for parsing the body.
// Is any error occurs it will panic.
// It's just a helper function to avoid writing if condition again n again.
func ParseBody(ctx *fiber.Ctx, body interface{}) error {
	if err := ctx.BodyParser(body); err != nil {
		return ErrInvalidArgument.WithMetadata(map[string]string{
			"error": err.Error(),
		})
	}
	return nil
}

// ParseBodyAndValidate is helper function for parsing the body.
// Is any error occurs it will panic.
// It's just a helper function to avoid writing if condition again n again.
func ParseBodyAndValidate(ctx *fiber.Ctx, body interface{}) error {
	if err := ParseBody(ctx, body); err != nil {
		return err
	}
	return Validate(body)
}

// ParseQuery ...
func ParseQuery(ctx *fiber.Ctx, query interface{}) error {
	if err := ctx.QueryParser(query); err != nil {
		return ErrInvalidArgument.WithMetadata(map[string]string{
			"error": err.Error(),
		})
	}
	return nil
}

// ParseQueryAndValidate ...
func ParseQueryAndValidate(ctx *fiber.Ctx, body interface{}) error {
	if err := ParseQuery(ctx, body); err != nil {
		return err
	}
	return Validate(body)
}

// Response is a API response
type Response struct {
	Code        int               `json:"code,omitempty"`         // 状态码
	Reason      string            `json:"reason,omitempty"`       // 错误原因
	Message     string            `json:"message,omitempty"`      // 错误信息，为用户可读的信息，可作为用户提示内容
	Data        interface{}       `json:"data,omitempty"`         // 正常的数据
	Metadata    map[string]string `json:"metadata,omitempty"`     // extra data
	RequestID   string            `json:"request_id,omitempty"`   //
	RequestTime int64             `json:"request_time,omitempty"` //
}

// ResOK 响应OK
func ResOK(ctx *fiber.Ctx) error {
	return ResSuccess(ctx, nil)
}

// ResList 响应列表数据
func ResList(ctx *fiber.Ctx, v interface{}) error {
	return ResSuccess(ctx, v)
}

// ResPage 响应分页数据
func ResPage(ctx *fiber.Ctx, v interface{}, pr *gormx.PaginationResult) error {
	list := gormx.ListResult{
		List:       v,
		Pagination: pr,
	}
	return ResSuccess(ctx, list)
}

// ResSuccess 响应成功
func ResSuccess(ctx *fiber.Ctx, v interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(Response{
		Code:        200,
		Message:     "请求成功",
		Data:        v,
		RequestID:   GetRequestID(ctx),
		RequestTime: time.Now().Unix(),
	})
}
