package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"fmt"
	"time"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	models "github.com/messagesgo/models"
)

func LoadEnv(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

const (
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)


func TranslateText(originalText string) (string, error) {
	apiKey := "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"
	client := resty.New()

	response, err := client.R().
		SetAuthToken(apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model":    "gpt-3.5-turbo",
			"messages": []interface{}{map[string]interface{}{"role": "system", "content": originalText +"translate into japanese"}},
			"max_tokens": 50,
		}).
		Post(apiEndpoint)

	if err != nil {
		return "", fmt.Errorf("Error while sending the translation request: %v", err)
	}

	body := response.Body()

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", fmt.Errorf("Error while decoding translation JSON response: %v", err)
	}

	// Extract the content from the JSON response
	translatedText := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	return translatedText, nil
}

func SaveMessage(message *models.Message) error {
	// Database connection parameters
	db := Connect()
	defer db.Close()

	// Prepare the SQL statement for inserting a message
	stmt, err := db.Prepare("INSERT INTO messages (japanese_text, english_text, user_id, company_id, chat_room_id, created_at) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("Error preparing SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the SQL statement with the message data
	_, err = stmt.Exec(message.JapaneseText, message.EnglishText, message.UserID, message.CompanyID, message.ChatRoomID, message.CreatedAt)
	if err != nil {
		return fmt.Errorf("Error executing SQL statement: %v", err)
	}

	return nil
}


func HandleSend(w http.ResponseWriter, r *http.Request) {
	originalText := r.FormValue("originalText")
	userID := r.FormValue("userID")
	companyID := r.FormValue("companyID")
	chatRoomID := r.FormValue("chatRoomID")

	translatedText, err := TranslateText(originalText)
	if err != nil {
		http.Error(w, "Error translating text", http.StatusInternalServerError)
		return
	}

	message := &models.Message{
		JapaneseText: translatedText,
		EnglishText:  originalText,
		UserID:       userID,
		CompanyID:    companyID,
		ChatRoomID:   chatRoomID,
		CreatedAt:    time.Now(),
	}

	err = SaveMessage(message)
	log.Println(err)
	if err != nil {
		http.Error(w, "Error saving message to the database", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Original Text: %s\nTranslated Text: %s\n", originalText, translatedText)

	// Respond with success or any necessary data
	response := map[string]string{
		"status":  "success",
		"message": "Message translated and saved successfully",
	}
	json.NewEncoder(w).Encode(response)
}
