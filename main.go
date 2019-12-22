package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imbaggaarm/go-messenger"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	kPort        string = "STRANGER_BOT_1_PORT"
	kVerifyToken string = "STRANGER_BOT_1_VERIFY_TOKEN"
	kAccessToken string = "STRANGER_BOT_1_ACCESS_TOKEN"
)

var bot *messenger.Bot

func main() {
	r := gin.Default()
	// Configure api sub router
	api := r.Group("/vnuchatbot/fb/v1")

	// Configure authentication endpoint
	api.GET("/webhook", VerificationEndpoint())
	// Configure message endpoint
	api.POST("/webhook", MessageEndpoint())
	// Configure get started button
	api.GET("/get_started", SetUpGetStarted())
	// Configure persistent menu
	api.GET("/persistent_menu", SetUpPersistentMenu())

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	port := os.Getenv(kPort)
	if port == "" {
		port = "8080"
	}

	accessToken := os.Getenv(kAccessToken)
	if accessToken == "" {
		log.Fatal("Failed to get page access token.\n")
	}

	apiVersion := messenger.DefaultApiVersion
	bot = messenger.NewBot(accessToken, apiVersion)

	if err := r.Run(":" + port); err != nil {
		log.Println("Failed to start server with error")
		log.Fatal(err)
	}
}

func VerificationEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		challenge := c.Query("hub.challenge")
		token := c.Query("hub.verify_token")
		if token == os.Getenv(kVerifyToken) {
			c.String(http.StatusOK, challenge)
		} else {
			c.Status(http.StatusForbidden)
		}
	}
}

func MessageEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Write response immediately
		c.String(http.StatusOK, "Message is processing")
		// Handle event
		var event messenger.WebhookEvent
		if err := c.BindJSON(&event); err != nil {
			log.Println(err.Error())
			return
		}
		go handleWebhookEvent(event)
	}
}

func SetUpGetStarted() gin.HandlerFunc {
	return func(c *gin.Context) {
		strIsOn := c.Query("is_on")
		isOn, err := strconv.ParseBool(strIsOn)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		if isOn {
			setGetStarted()
		} else {
			removeGetStarted()
		}
	}
}

func SetUpPersistentMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		strIsOn := c.Query("is_on")
		isOn, err := strconv.ParseBool(strIsOn)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		if isOn {
			setPersistentMenu()
		} else {
			removeGetStarted()
		}
	}
}
