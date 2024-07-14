package custom

import (
	"github.com/gofiber/fiber/v3"
)

type (
	SuccessJSON struct {
		StatusCode int    `json:"statusCode"`
		Status     string `json:"status"`
		Message    string `json:"message"`
	}

	DataJSON struct {
		StatusCode int    `json:"statusCode"`
		Status     string `json:"status"`
		Data       any    `json:"data"`
	}

	ErrorJSON struct {
		StatusCode int    `json:"statusCode"`
		Status     string `json:"status"`
		Error      string `json:"error"`
	}
)

func SuccessResponse(ctx fiber.Ctx, statusCode int, message string) error {
	return ctx.Status(statusCode).JSON(&SuccessJSON{
		StatusCode: statusCode,
		Status:     MapStatusCode(statusCode),
		Message:    message,
	})
}

func DataResponse(ctx fiber.Ctx, statusCode int, data any) error {
	return ctx.Status(statusCode).JSON(&DataJSON{
		StatusCode: statusCode,
		Status:     MapStatusCode(statusCode),
		Data:       data,
	})
}

func ErrorResponse(ctx fiber.Ctx, fiberError *fiber.Error, err error) error {
	return ctx.Status(fiberError.Code).JSON(&ErrorJSON{
		StatusCode: fiberError.Code,
		Status:     fiberError.Message,
		Error:      err.Error(),
	})
}

func MapStatusCode(statusCode int) string {
	var status string

	switch statusCode {
	case 200:
		status = "OK"
	case 201:
		status = "Created"
	case 202:
		status = "Accepted"
	case 204:
		status = "No Content"
	default:
		status = "OK"
	}

	return status
}
