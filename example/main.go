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
