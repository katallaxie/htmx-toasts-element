# 🥪 HTMX Toasts Element

[![Node CI](https://github.com/katallaxie/htmx-toasts-element/actions/workflows/main.yml/badge.svg)](https://github.com/katallaxie/htmx-toasts-element/actions/workflows/main.yml)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)
[![Volkswagen](https://auchenberg.github.io/volkswagen/volkswargen_ci.svg?v=1)](https://github.com/auchenberg/volkswagen)
![GitHub License](https://img.shields.io/github/license/katallaxie/htmx-toasts-element)

This is a simple HTMX element that can be used to display toasts in your web application.

## Usage

```html
<template>
  <div class="alert" slot="alert">
    <span slot="message"></span>
    <button type="button" class="btn btn-sm btn-outline" aria-label="Close" slot="close">Close</button>
  </div>
</template>
<htmx-toasts
  timeout="3000"
  class="toast"
  role="status"
  aria-live="polite"
  success-class="alert-success"
  error-class="alert-error"
  info-class="alert-info"
  warn-class="alert-warning"></htmx-toasts>
```

## Attributes

- `timeout` - The time in milliseconds that the toast will be displayed.
- `success-class` - The class to apply to the toast when it is a success.
- `error-class` - The class to apply to the toast when it is an error.
- `info-class` - The class to apply to the toast when it is an info message.
- `warn-class` - The class to apply to the toast when it is a warning.
- `template-id`- The id of the `template` element to use to insert in the toast.
- `stylesheet-href` - The url to the stylesheet for the toasts.

## Installation

```bash
npm install @htmx/htmx-toast
```

Use as a module in your application:

```html
<script src="https://unpkg.com/@htmx/htmx-toasts@latest/dist/index.js" type="module"></script>
```

## License

[MIT](/LICENSE)