package main

import (
	"fmt"
	"hello_world/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize HTTP server
	router := gin.Default()

	router.GET("/xxxx-xxxx/v1/message", func(c *gin.Context) {

		log.Println("Received /xxxx-xxxx/v1/message request .")

		for i := 0; i < 100; i++ {
			logger.LoggerS3.Println("Hello Jamesbond from /xxxx-xxxx/v1/message logger at", time.Now(), " -> ", i)
		}

		c.JSON(http.StatusOK, "Hello ....xxxx...")
	})

	router.GET("/yyyy-yyyy/v1/message-y", func(c *gin.Context) {

		log.Println("Received /yyyy-yyyy/v1/message request .")

		c.JSON(http.StatusOK, "Hello ____yyyy____")
	})

	// Start HTTP server
	go func() {
		if err := router.Run(":8080"); err != nil {
			fmt.Printf("Error starting HTTP server: %v\n", err)
			os.Exit(1)
		}
	}()

	// router for healthz endpoints.
	healthzRouter := gin.Default()
	go func() {
		err := healthzRouter.Run(":5051")
		if err != nil {
			log.Fatalf("Unable to run gin healthz server %v.", err)
		}
	}()

	// Graceful shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	<-sigchan

	fmt.Println("Shutting down gracefully...")
}
