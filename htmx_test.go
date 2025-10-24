package toasts_test

import (
	"bytes"
	"strings"
	"testing"

	htmx "github.com/katallaxie/htmx"
	toasts "github.com/katallaxie/htmx-toasts-element"
	"github.com/stretchr/testify/require"
)

func TestHTMXToasts(t *testing.T) {
	tests := []struct {
		name string
		node htmx.Node
		want string
	}{
		{
			name: "HTMXToasts",
			node: toasts.Toasts(),
			want: `<htmx-toasts></htmx-toasts>`,
		},
		{
			name: "HTMXToastsTimeout",
			node: toasts.HTMXToastsTimeout(5000),
			want: `timeout="5000"`,
		},
		{
			name: "HTMXToastsErrorClass",
			node: toasts.HTMXToastsErrorClass("error-class"),
			want: `error-class="error-class"`,
		},
		{
			name: "HTMXToastsSuccessClass",
			node: toasts.HTMXToastsSuccessClass("success-class"),
			want: `success-class="success-class"`,
		},
		{
			name: "HTMXToastsInfoClass",
			node: toasts.HTMXToastsInfoClass("info-class"),
			want: `info-class="info-class"`,
		},
		{
			name: "HTMXToastsWarningClass",
			node: toasts.HTMXToastsWarningClass("warning-class"),
			want: `warning-class="warning-class"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := tt.node.Render(&buf)
			require.NoError(t, err)
			require.Equal(t, tt.want, strings.TrimSpace(buf.String()))
		})
	}
}
