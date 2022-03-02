# rollbar

[![Run Tests](https://github.com/gin-contrib/rollbar/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/rollbar/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/rollbar/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/rollbar)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-contrib/rollbar)](https://goreportcard.com/report/github.com/gin-contrib/rollbar)
[![GoDoc](https://godoc.org/github.com/gin-contrib/rollbar?status.svg)](https://godoc.org/github.com/gin-contrib/rollbar)
[![Join the chat at https://gitter.im/gin-gonic/gin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

Middleware to integrate with [rollbar](https://rollbar.com/) error monitoring. It uses [rollbar-go](https://github.com/rollbar/rollbar-go) SDK that reports errors and logs messages.

## Usage

Download and install it:

```sh
go get github.com/ForwardFinancing/ff-rollbar
```

Import it in your code:

```go
import "github.com/ForwardFinancing/ff-rollbar"
```

## Example

```go
package main

import (
  "log"

  ffrollbar "github.com/ForwardFinancing/ff-rollbar"
  "github.com/gin-gonic/gin"

  "github.com/rollbar/rollbar-go"
)

func main() {
  rollbar.SetToken("MY_TOKEN")
  // rollbar.SetEnvironment("production") // defaults to "development"

  r := gin.Default()
  r.Use(ffrollbar.Recovery(true))

  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```
