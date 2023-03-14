package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

var (
	errMissingTicker = errors.New("Missing ticker param.")
)

func errorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"data":  "",
		"error": err.Error(),
	}
}

func successResponse(arg interface{}) *fiber.Map {
	return &fiber.Map{
		"data":  arg,
		"error": nil,
	}
}
