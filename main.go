package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"

	"github.com/labstack/echo/v4/middleware"
)

type Message struct {
	UserID string `json:"userId"`
	Text   string `json:"text"`
}

type CardAttachment struct {
	Title       string `json:"title"`
	Text        string `json:"text"`
	Markdown    bool   `json:"markdown"`
	ContentType string `json:"contentType"`
}

type CardMessage struct {
	Type        string           `json:"type"`
	Attachments []CardAttachment `json:"attachments"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	// Endpoint to handle incoming messages
	e.POST("/api/messages", handleMessage)

	// Start the server
	e.Start(":8080")
}

func handleMessage(c echo.Context) error {
	// Parse the incoming message
	var message Message
	if err := c.Bind(&message); err != nil {
		log.Println("Error parsing message:", err)
		return c.NoContent(http.StatusBadRequest)
	}

	// Process the message and generate the card
	cardAttachment := generateCard(message.Text)

	// Prepare the card message
	cardMessage := CardMessage{
		Type:        "message",
		Attachments: []CardAttachment{cardAttachment},
	}

	// Send the card as the response
	return c.JSON(http.StatusOK, cardMessage)
}

func generateCard(update string) CardAttachment {
	// Process the user's daily update

	// Here, you can implement your logic to handle the user's daily update
	// You can parse the text and extract the necessary information
	// You can store the updates in a database or perform any other actions based on your requirements

	// Generate a card with the user's update
	cardAttachment := CardAttachment{
		Title:       "Daily Update",
		Text:        update,
		Markdown:    true,
		ContentType: "application/vnd.microsoft.card.adaptive",
	}

	return cardAttachment
}
