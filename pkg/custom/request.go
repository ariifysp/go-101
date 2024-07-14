package custom

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type (
	FiberRequestInterface interface {
		BindQuery(obj any) error
		BindBody(obj any) error
	}

	CustomFiberRequest struct {
		ctx       fiber.Ctx
		validator *validator.Validate
	}
)

var (
	once              sync.Once
	validatorInstance *validator.Validate
)

func NewCustomFiberRequest(fiberRequest fiber.Ctx) FiberRequestInterface {
	once.Do(func() {
		validatorInstance = validator.New()
	})

	return &CustomFiberRequest{
		ctx:       fiberRequest,
		validator: validatorInstance,
	}
}

func (r *CustomFiberRequest) BindQuery(obj any) error {
	if err := r.ctx.Bind().Query(obj); err != nil {
		return err
	}

	if err := r.validator.Struct(obj); err != nil {
		return err
	}

	return nil
}

func (r *CustomFiberRequest) BindBody(obj any) error {
	if err := r.ctx.Bind().Body(obj); err != nil {
		return err
	}

	if err := r.validator.Struct(obj); err != nil {
		return err
	}

	return nil
}
