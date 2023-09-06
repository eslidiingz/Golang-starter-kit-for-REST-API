package controllers

import (
	"api-service-sdk/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LivestremControllerStruct struct{}

func LivestremController() *LivestremControllerStruct {
	return &LivestremControllerStruct{}
}

// Define the payload structure
type Payload struct {
	Meta      map[string]string `json:"meta"`
	Recording map[string]string `json:"recording"`
}

type RtmpsStruct struct {
	Url       string `json:"url"`
	StreamKey string `json:"streamKey"`
}

type ResultStruct struct {
	ResultUid   string      `json:"uid"`
	ResultRtmps RtmpsStruct `json:"rtmps"`
}

type GenerateKeyResult struct {
	GenerateResult  ResultStruct `json:"result"`
	GenerateSuccess bool         `json:"success"`
}

func GetLiveStream(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "livestream key for user authenticated ^^",
	})
}

type LivestreamRequest struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type LiveStreamCreatedResult struct {
	StreamKey string `json:"streamKey"`
	StreamUrl string `json:"streamUrl"`
	M3u8Url   string `json:"m3u8Url"`
}

func CreateLivestream(db *gorm.DB, c *gin.Context) {
	// user, _ := c.Get("user")

	var livestreamReq LivestreamRequest

	if err := c.ShouldBindJSON(&livestreamReq); err != nil {
		c.JSON(400, gin.H{"error": "Missing or invalid fields"})
		return
	}

	genKeyResult := generateStreamingKey(livestreamReq.Name)
	result := genKeyResult.GenerateResult

	streamKey := result.ResultRtmps.StreamKey
	streamUrl := result.ResultRtmps.Url
	m3u8Url := "https://" + os.Getenv("CLOUDFLARE_CUSTOMER_SUBDOMAIN") + "/" + result.ResultUid + "/manifest/video.m3u8"

	// Cconvert json to string
	var responseResult GenerateKeyResult
	responseResult.GenerateResult = result
	out, _ := json.Marshal(&responseResult)

	// Assign the data to model
	var livestream models.Livestream
	livestream.Name = livestreamReq.Name
	livestream.StreamKey = streamKey
	livestream.StreamUrl = streamUrl
	livestream.M3u8Url = m3u8Url
	livestream.Response = string(out)

	// Bind the JSON body to the Livestream struct
	c.BindJSON(&livestream)

	// Save to database
	created := db.Create(&livestream)

	if created.Error != nil {
		c.JSON(500, gin.H{
			"error":   true,
			"message": "Could not create livestream",
		})
		return
	}

	// Return response
	c.JSON(200, gin.H{
		"success": true,
		"message": "Livestream has been created success.",
		"result":  livestream,
	})
}

func generateStreamingKey(_name string) GenerateKeyResult {
	providerApi := os.Getenv("CLOUDFLARE_URI")
	client := &http.Client{}

	// Create the payload
	payloadData := Payload{
		Meta:      map[string]string{"name": _name},
		Recording: map[string]string{"mode": "automatic"},
	}

	// Convert payload to JSON
	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		log.Fatalf("Could not marshal payload: %s", err)
	}

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", providerApi, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Failed to create request: %s", err)
	}

	// Add headers to the request
	req.Header.Set("Authorization", "Bearer "+os.Getenv("CLOUDFLARE_API_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request failed: %s", err)
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %s", err)
	}

	// log.Printf("Response: %s", body)

	var responseResult GenerateKeyResult
	if err := json.Unmarshal([]byte(body), &responseResult); err != nil {
		fmt.Println("Error unmarshalling:", err)
		// return
	}
	// fmt.Println("Unmarshalled data:", responseResult)

	return responseResult
}
