package ffrollbar

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/rollbar/rollbar-go"
)

// Recovery middleware for rollbar error monitoring
func Recovery(token string, environment string, onlyCrashes bool) gin.HandlerFunc {
	// Recover without logging to Rollbar if no token is provided
	if token == "" {
		return func(c *gin.Context) {
			defer func() {
				if rval := recover(); rval != nil {
					debug.PrintStack()
					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}()

			c.Next()
		}
	}

	// Configure since token is provided
	rollbar.SetToken(token)
	rollbar.SetEnvironment(translateEnvironment(environment))

	return func(c *gin.Context) {
		defer func() {
			if rval := recover(); rval != nil {
				debug.PrintStack()

				rollbar.Critical(errors.New(fmt.Sprint(rval)), map[string]interface{}{
					"endpoint": c.Request.RequestURI,
					"params":   c.Request.URL.Query(),
				})

				c.AbortWithStatus(http.StatusInternalServerError)
			}

			if !onlyCrashes {
				for _, item := range c.Errors {
					rollbar.Error(item.Err, map[string]interface{}{
						"meta":     fmt.Sprint(item.Meta),
						"endpoint": c.Request.RequestURI,
						"params":   c.Request.URL.Query(),
					})
				}
			}
		}()

		c.Next()
	}
}

func translateEnvironment(environment string) string {
	if environment == "prod" {
		return "production"
	} else {
		return environment
	}
}
