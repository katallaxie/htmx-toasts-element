package toasts

import (
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v3"
	handler "github.com/katallaxie/fiber-htmx/v3"
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/pkg/conv"
)

// ErrToastsUnexpectedError is returned when there is an unexpected error.
var ErrToastsUnexpectedError = New(ERROR, "An unexpected error occurred", fiber.StatusInternalServerError)

const (
	// INFO is the info level.
	INFO = "info"
	// SUCCESS is the success level.
	SUCCESS = "success"
	// ERROR is the error level.
	ERROR = "error"
)

// Errors returns a Fiber error handler that formats errors into a consistent JSON structure.
func Errors() fiber.ErrorHandler {
	return func(c fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		toast := ErrToastsUnexpectedError

		var e *fiber.Error // if this is not a toast then use the error message
		if errors.As(err, &e) {
			code = e.Code
			toast = Error(e.Message)
		}

		if errors.As(err, &toast) {
			code = toast.Code
		}

		if toast.Level != SUCCESS {
			handler.ReSwap(c, "none")
		}

		if err := toast.SetHXTriggerHeader(c); err != nil {
			return err
		}

		c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

		return c.Status(code).SendString(err.Error())
	}
}

// Toast is a message to display to the user.
type Toast struct {
	// Level is the level of the toast.
	Level string `json:"level"`
	// Message is the message of the toast.
	Message string `json:"message"`
	// Code is the http status code of the toast.
	Code int `json:"code"`
}

// New returns a new Toast.
func New(level, message string, code ...int) *Toast {
	statusCode := 200
	if len(code) > 0 {
		statusCode = code[0]
	}

	return &Toast{level, message, statusCode}
}

// Error returns the error message.
func (t *Toast) Error() string {
	return t.Message
}

// Info returns an info message.
func Info(message string, code ...int) *Toast {
	statusCode := 200
	if len(code) > 0 {
		statusCode = code[0]
	}

	return New(INFO, message, statusCode)
}

// Success returns a success message.
func Success(c fiber.Ctx, message string, code ...int) {
	statusCode := 200
	if len(code) > 0 {
		statusCode = code[0]
	}

	//nolint:errcheck
	New(SUCCESS, message, statusCode).SetHXTriggerHeader(c)
}

// Error returns an error message.
func Error(message string, code ...int) *Toast {
	statusCode := 500
	if len(code) > 0 {
		statusCode = code[0]
	}

	return &Toast{
		Level:   ERROR,
		Message: message,
		Code:    statusCode,
	}
}

// ToJSON returns the JSON representation of the toast.
func (t Toast) ToJSON() (string, error) {
	t.Message = t.Error()

	eventMap := map[string]Toast{}
	eventMap["htmx-toasts:notify"] = t

	jsonData, err := json.Marshal(eventMap)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// SetHXTriggerHeader sets the HTMX trigger header.
func (t Toast) SetHXTriggerHeader(c fiber.Ctx) error {
	jsonData, err := t.ToJSON()
	if err != nil {
		return err
	}

	handler.Trigger(c, jsonData)

	return nil
}

// Toast is the toast component.
func Toasts() htmx.Node {
	return htmx.CustomElement("htmx-toasts")
}

// HTMXToastsTimeout sets the timeout for the toasts in milliseconds.
func HTMXToastsTimeout(timeout int) htmx.Node {
	return htmx.Attribute("timeout", conv.String(timeout))
}

// HTMXToastsErrorClass sets the CSS class for error toasts.
func HTMXToastsErrorClass(class string) htmx.Node {
	return htmx.Attribute("error-class", class)
}

// HTMXToastsSuccessClass sets the CSS class for success toasts.
func HTMXToastsSuccessClass(class string) htmx.Node {
	return htmx.Attribute("success-class", class)
}

// HTMXToastsInfoClass sets the CSS class for info toasts.
func HTMXToastsInfoClass(class string) htmx.Node {
	return htmx.Attribute("info-class", class)
}

// HTMXToastsWarningClass sets the CSS class for warning toasts.
func HTMXToastsWarningClass(class string) htmx.Node {
	return htmx.Attribute("warning-class", class)
}

// HTMXToastsStyleSheetHref returns the href for the HTMX toasts stylesheet.
func HTMXToastsStyleSheetHref(href string) htmx.Node {
	return htmx.Attribute("stylesheet-href", href)
}

// HTMXToastsTemplateID sets the template ID for the toasts.
func HTMXToastsTemplateID(id string) htmx.Node {
	return htmx.Attribute("template-id", id)
}
