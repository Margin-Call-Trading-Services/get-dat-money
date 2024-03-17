package model

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrMissingTicker = errors.New("missing ticker param")
)

func ErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"data":  "",
		"error": err.Error(),
	}
}

func SuccessResponse(arg interface{}) *fiber.Map {
	return &fiber.Map{
		"data":  arg,
		"error": nil,
	}
}
