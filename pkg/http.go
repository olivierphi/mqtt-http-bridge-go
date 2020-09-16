package pkg

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartHttpServer(port uint16) {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.Match([]string{"GET", "HEAD"}, "/ping", ping)
	e.POST("/api/mqtt/msg", publishMqttMessage)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

func ping(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @link https://echo.labstack.com/guide/request#validate-data
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

type PublishMqttMessageRequestPayload struct {
	Topic      string `json:"topic" form:"topic" validate:"required"`
	Message    string `json:"message" form:"message" validate:"required"`
	Qos        byte   `json:"qos" form:"qos" validate:"required"`
	IsRetained bool   `json:"isRetained" form:"isRetained"`
}

func publishMqttMessage(c echo.Context) (err error) {
	payload := new(PublishMqttMessageRequestPayload)
	if err = c.Bind(payload); err != nil {
		return
	}
	fmt.Printf("payload: %+v", payload)
	if err = c.Validate(payload); err != nil {
		fmt.Printf("validation error: %v", err)
		return
	}
	return c.NoContent(http.StatusOK)
}
