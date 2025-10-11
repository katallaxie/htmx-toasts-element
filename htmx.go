package toasts

import (
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/pkg/conv"
)

// HTMXToasts is a component that displays toasts.
func HTMXToasts() htmx.Node {
	return htmx.Element("htmx-toasts")
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
